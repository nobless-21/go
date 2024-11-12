//nolint:typecheck
package main

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

const (
	// какой адрес-порт слушать серверу
	listenAddr string = "127.0.0.1:8082"

	// кого по каким методам пускать
	ACLData string = `{
	"logger":    ["/main.Admin/Logging"],
	"stat":      ["/main.Admin/Statistics"],
	"biz_user":  ["/main.Biz/Check", "/main.Biz/Add"],
	"biz_admin": ["/main.Biz/*"]
}`
)

// чтобы не было сюрпризов когда где-то не успела переключиться горутина и не успело что-то стартовать
func wait(amount int) {
	time.Sleep(time.Duration(amount) * 10 * time.Millisecond)
}

// утилитарная функция для коннекта к серверу
func getGrpcConn(t *testing.T) *grpc.ClientConn {
	grpcConn, err := grpc.Dial(
		listenAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("cant connect to grpc: %v", err)
	}
	return grpcConn
}

// получаем контекст с нужными метаданными для ACL
func getConsumerCtx(consumerName string) context.Context {
	// ctx, _ := context.WithTimeout(context.Background(), time.Second)
	ctx := context.Background()
	md := metadata.Pairs(
		"consumer", consumerName,
	)
	return metadata.NewOutgoingContext(ctx, md)
}

// старт-стоп сервера
func TestServerStartStop(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	wait(1)
	finish() // при вызове этой функции ваш сервер должен остановиться и освободить порт
	wait(1)

	// теперь проверим что вы освободили порт и мы можем стартовать сервер ещё раз
	ctx, finish = context.WithCancel(context.Background())
	err = StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		t.Fatalf("cant start server again: %v", err)
	}
	wait(1)
	finish()
	wait(1)
}

// у вас наверняка будет что-то выполняться в отдельных горутинах
// этим тестом мы проверяем что вы останавливаете все горутины которые у вас были и нет утечек
// некоторый запас ( goroutinesPerTwoIterations*5 ) остаётся на случай рантайм горутин
func TestServerLeak(t *testing.T) {
	// return
	goroutinesStart := runtime.NumGoroutine()
	TestServerStartStop(t)
	goroutinesPerTwoIterations := runtime.NumGoroutine() - goroutinesStart

	goroutinesStart = runtime.NumGoroutine()
	goroutinesStat := []int{}
	for i := 0; i <= 25; i++ {
		TestServerStartStop(t)
		goroutinesStat = append(goroutinesStat, runtime.NumGoroutine())
	}
	goroutinesPerFiftyIterations := runtime.NumGoroutine() - goroutinesStart
	if goroutinesPerFiftyIterations > goroutinesPerTwoIterations*5 {
		t.Fatalf("looks like you have goroutines leak: %+v", goroutinesStat)
	}
}

// ACL (права на методы доступа) парсится корректно
func TestACLParseError(t *testing.T) {
	// finish'а тут нет потому что стартовать у вас ничего не должно если не получилось распаковать ACL
	err := StartMyMicroservice(context.Background(), listenAddr, "{.;")
	if err == nil {
		t.Fatalf("expacted error on bad acl json, have nil")
	}
}

// ACL (права на методы доступа) работает корректно
func TestACL(t *testing.T) {
	wait(1)
	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	wait(1)
	defer func() {
		finish()
		wait(1)
	}()

	conn := getGrpcConn(t)
	defer conn.Close()

	biz := NewBizClient(conn)
	adm := NewAdminClient(conn)

	for idx, ctx := range []context.Context{
		context.Background(),       // нет поля для ACL
		getConsumerCtx("unknown"),  // поле есть, неизвестный консюмер
		getConsumerCtx("biz_user"), // поле есть, нет доступа
	} {
		_, err = biz.Test(ctx, &Nothing{})
		if err == nil {
			t.Fatalf("[%d] ACL fail: expected err on disallowed method", idx)
		} else if code := status.Code(err); code != codes.Unauthenticated {
			t.Fatalf("[%d] ACL fail: expected Unauthenticated code, got %v", idx, code)
		}
	}

	// есть доступ
	_, err = biz.Check(getConsumerCtx("biz_user"), &Nothing{})
	if err != nil {
		t.Fatalf("ACL fail: unexpected error: %v", err)
	}
	_, err = biz.Check(getConsumerCtx("biz_admin"), &Nothing{})
	if err != nil {
		t.Fatalf("ACL fail: unexpected error: %v", err)
	}
	_, err = biz.Test(getConsumerCtx("biz_admin"), &Nothing{})
	if err != nil {
		t.Fatalf("ACL fail: unexpected error: %v", err)
	}

	// ACL на методах, которые возвращают поток данных
	logger, err := adm.Logging(getConsumerCtx("unknown"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = logger.Recv()
	if err == nil {
		t.Fatalf("ACL fail: expected err on disallowed method")
	} else if code := status.Code(err); code != codes.Unauthenticated {
		t.Fatalf("ACL fail: expected Unauthenticated code, got %v", code)
	}
}

func TestLogging(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	wait(1)
	defer func() {
		finish()
		wait(1)
	}()

	conn := getGrpcConn(t)
	defer conn.Close()

	biz := NewBizClient(conn)
	adm := NewAdminClient(conn)

	logStream1, err := adm.Logging(getConsumerCtx("logger"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1 * time.Millisecond)

	logStream2, err := adm.Logging(getConsumerCtx("logger"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}

	logData1 := []*Event{}
	logData2 := []*Event{}

	wait(1)

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-time.After(3 * time.Second):
			fmt.Println("looks like you dont send anything to log stream in 3 sec")
			t.Errorf("looks like you dont send anything to log stream in 3 sec")
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 4; i++ {
			evt, errTmp := logStream1.Recv()
			// log.Println("logger 1", evt, errTmp)
			if errTmp != nil {
				t.Errorf("unexpected error: %v, awaiting event", errTmp)
				return
			}
			if !strings.HasPrefix(evt.GetHost(), "127.0.0.1:") || evt.GetHost() == listenAddr {
				t.Errorf("bad host: %v", evt.GetHost())
				return
			}
			// это грязный хак
			// protobuf добавляет к структуре свои поля, которвые не видны при приведении к строке и при reflect.DeepEqual
			// поэтому берем не оригинал сообщения, а только нужные значения
			logData1 = append(logData1, &Event{Consumer: evt.Consumer, Method: evt.Method})
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			evt, errTmp := logStream2.Recv()
			// log.Println("logger 2", evt, errTmp)
			if errTmp != nil {
				t.Errorf("unexpected error: %v, awaiting event", errTmp)
				return
			}
			if !strings.HasPrefix(evt.GetHost(), "127.0.0.1:") || evt.GetHost() == listenAddr {
				t.Errorf("bad host: %v", evt.GetHost())
				return
			}
			// это грязный хак
			// protobuf добавляет к структуре свои поля, которвые не видны при приведении к строке и при reflect.DeepEqual
			// поэтому берем не оригинал сообщения, а только нужные значения
			logData2 = append(logData2, &Event{Consumer: evt.Consumer, Method: evt.Method})
		}
	}()

	_, err = biz.Check(getConsumerCtx("biz_user"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Millisecond)

	_, err = biz.Check(getConsumerCtx("biz_admin"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Millisecond)

	_, err = biz.Test(getConsumerCtx("biz_admin"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Millisecond)

	wg.Wait()

	expectedLogData1 := []*Event{
		{Consumer: "logger", Method: "/main.Admin/Logging"},
		{Consumer: "biz_user", Method: "/main.Biz/Check"},
		{Consumer: "biz_admin", Method: "/main.Biz/Check"},
		{Consumer: "biz_admin", Method: "/main.Biz/Test"},
	}
	expectedLogData2 := []*Event{
		{Consumer: "biz_user", Method: "/main.Biz/Check"},
		{Consumer: "biz_admin", Method: "/main.Biz/Check"},
		{Consumer: "biz_admin", Method: "/main.Biz/Test"},
	}

	if !reflect.DeepEqual(logData1, expectedLogData1) {
		t.Fatalf("logs1 dont match\nhave %+v\nwant %+v", logData1, expectedLogData1)
	}
	if !reflect.DeepEqual(logData2, expectedLogData2) {
		t.Fatalf("logs2 dont match\nhave %+v\nwant %+v", logData2, expectedLogData2)
	}
}

func TestStat(t *testing.T) {
	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		t.Fatalf("cant start server initial: %v", err)
	}
	wait(1)
	defer func() {
		finish()
		wait(2)
	}()

	conn := getGrpcConn(t)
	defer conn.Close()

	biz := NewBizClient(conn)
	adm := NewAdminClient(conn)

	statStream1, err := adm.Statistics(getConsumerCtx("stat"), &StatInterval{IntervalSeconds: 2})
	if err != nil {
		t.Fatal(err)
	}
	wait(1)
	statStream2, err := adm.Statistics(getConsumerCtx("stat"), &StatInterval{IntervalSeconds: 3})
	if err != nil {
		t.Fatal(err)
	}

	mu := &sync.Mutex{}
	stat1 := &Stat{}
	stat2 := &Stat{}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			stat, errTmp := statStream1.Recv()
			if errTmp != nil && errTmp != io.EOF {
				// fmt.Printf("unexpected error %v\n", errTmp)
				return
			} else if errTmp == io.EOF {
				break
			}
			// log.Println("stat1", stat, errTmp)
			mu.Lock()
			// это грязный хак
			// protobuf добавляет к структуре свои поля, которвые не видны при приведении к строке и при reflect.DeepEqual
			// поэтому берем не оригинал сообщения, а только нужные значения
			stat1 = &Stat{
				ByMethod:   stat.ByMethod,
				ByConsumer: stat.ByConsumer,
			}
			mu.Unlock()
		}
	}()
	go func() {
		for {
			stat, errTmp := statStream2.Recv()
			if errTmp != nil && errTmp != io.EOF {
				// fmt.Printf("unexpected error %v\n", errTmp)
				return
			} else if errTmp == io.EOF {
				break
			}
			// log.Println("stat2", stat, errTmp)
			mu.Lock()
			// это грязный хак
			// protobuf добавляет к структуре свои поля, которвые не видны при приведении к строке и при reflect.DeepEqual
			// поэтому берем не оригинал сообщения, а только нужные значения
			stat2 = &Stat{
				ByMethod:   stat.ByMethod,
				ByConsumer: stat.ByConsumer,
			}
			mu.Unlock()
		}
	}()

	wait(1)

	_, err = biz.Check(getConsumerCtx("biz_user"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = biz.Add(getConsumerCtx("biz_user"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}
	_, err = biz.Test(getConsumerCtx("biz_admin"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}

	wait(200) // 2 sec

	expectedStat1 := &Stat{
		ByMethod: map[string]uint64{
			"/main.Biz/Check":        1,
			"/main.Biz/Add":          1,
			"/main.Biz/Test":         1,
			"/main.Admin/Statistics": 1,
		},
		ByConsumer: map[string]uint64{
			"biz_user":  2,
			"biz_admin": 1,
			"stat":      1,
		},
	}

	mu.Lock()
	if !reflect.DeepEqual(stat1, expectedStat1) {
		t.Fatalf("stat1-1 dont match\nhave %+v\nwant %+v", stat1, expectedStat1)
	}
	mu.Unlock()

	_, err = biz.Add(getConsumerCtx("biz_admin"), &Nothing{})
	if err != nil {
		t.Fatal(err)
	}

	wait(220) // 2+ sec

	expectedStat1 = &Stat{
		Timestamp: 0,
		ByMethod: map[string]uint64{
			"/main.Biz/Add": 1,
		},
		ByConsumer: map[string]uint64{
			"biz_admin": 1,
		},
	}
	expectedStat2 := &Stat{
		Timestamp: 0,
		ByMethod: map[string]uint64{
			"/main.Biz/Check": 1,
			"/main.Biz/Add":   2,
			"/main.Biz/Test":  1,
		},
		ByConsumer: map[string]uint64{
			"biz_user":  2,
			"biz_admin": 2,
		},
	}

	mu.Lock()
	if !reflect.DeepEqual(stat1, expectedStat1) {
		t.Fatalf("stat1-2 dont match\nhave %+v\nwant %+v", stat1, expectedStat1)
	}
	if !reflect.DeepEqual(stat2, expectedStat2) {
		t.Fatalf("stat2 dont match\nhave %+v\nwant %+v", stat2, expectedStat2)
	}
	mu.Unlock()

	finish()
}
