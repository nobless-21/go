package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type kitchen_objects struct {
	status int
	tea int
}

type hallway struct {
	status int
}

type room struct {
	key int
	status int
	synopsis int 
	backpack int 
}

type street struct{
	status int
}

type inventory struct {
	backpack_status int
	key int
	synopsis int 
	backpack int 
}

type game struct{
	Kitchen kitchen_objects
	hallway hallway
	room room
	street street
	inventory inventory
}

var status int = 0

func main() {
	fmt.Println("Выбери команду: осмотреться")
	status = 1
	initGame()
}

func initGame() {

	take := make(map[string]string)
	take["осмотреться"] = "Take_look_around"
	take["взять"] = "take"
	take["надеть"] = "put_on"
	take["идти"] = "walk"

	var game game

	game.Kitchen.status = 1
	game.Kitchen.tea = 1
	game.hallway.status = 0
	game.room.status = 0
	game.room.key = 1
	game.room.synopsis = 1
	game.room.backpack = 1
	game.street.status = 0
	game.inventory.backpack_status = 0
	game.inventory.key = 0
	game.inventory.synopsis = 0
	game.inventory.backpack = 0


	for status == 1 {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		commandFromUser := strings.TrimSuffix(text, "\n")
		words := strings.Fields(commandFromUser)
		command, ok := take[words[0]]
		if ok {
			result := handleCommand(commandFromUser, &game, command)
			fmt.Println(result)
			if result == "на улице весна. можно пройти - домой" {
				status = 0
			}
		}else{
			fmt.Println("Команда введена не коректно!")
		}
	}
}

func handleCommand(commandFromUser string, game *game, command string) string {
	keys := make(map[string]string)
	
	keys["идти коридор"] = "Walk_hallway"
	keys["идти комната"] = "Walk_room"
	keys["идти улица"] = "Walk_street"
	keys["идти кухня"] = "Walk_kitchen"
	keys["идти улица"] = "Walk_street"
	keys["взять ключи"] = "Take_key"
	keys["взять конспекты"] = "Take_synopsis"
	keys["надеть рюкзак"] = "Put_on_backpack"
	keys["осмотреться"] = "Take_look_around"

	commandKey, ok := keys[commandFromUser]
	if ok{ 
		if command == "walk"{
			switch commandKey {
				case "Walk_hallway":
					return Walk_hallway(game)
				case "Walk_room":
					return Walk_room(game)
				case "Walk_kitchen":
					return Walk_kitchen(game)
				case "Walk_street":
					return Walk_street(game)
				default:
					return "неизвестная команда"
			}
		}else if commandKey == "Take_look_around" {
			return Take_look_around(*game)
		}else if command == "take"{
			switch commandKey {
				case "Take_key":
					return Take_key(game)
				case "Take_synopsis":
					return Take_synopsis(game)
				default:
					return "неизвестная команда"
			}
		}else if command == "put_on"{
			switch commandKey {
				case "Put_on_backpack":
					return Put_on_backpack(game)
				default:
					return "неизвестная команда"
				}
		}else{
			return "неизвестная команда,"
		}
	}else{
		if command == "walk"{
			return "неизвестная команда"
		}else if commandKey == "Take_look_around" {
			return Take_look_around(*game)
		}else{
			return "нет такого"
		}
	}
	
}

 func Take_look_around(game game) string {
	var result string
	if game.Kitchen.status == 1 {
		if game.Kitchen.tea == 1 {
			result = "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		} else {
			result =  "ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		}
	} else if game.hallway.status == 1{
		result = "ты находишьс в коридоре, можно пройти на кухню или в комнату"
	} else if game.room.status == 1{
		if game.room.synopsis == 1 && game.room.backpack == 1 && game.room.key == 1 {
			result = "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"
		}
		if game.room.synopsis == 1 && game.room.backpack == 1 && game.room.key == 0 {
			result = "на столе: конспекты, на стуле: рюкзак, можно пройти - коридор"
		}
		if game.room.synopsis == 1 && game.room.backpack == 0 && game.room.key == 1 {
			result = "на столе: конспекты, ключи, можно пройти - коридор"
		}
		if game.room.synopsis == 1 && game.room.backpack == 0 && game.room.key == 0 {
			result = "на столе: конспекты, можно пройти - коридор"
		}
		if game.room.synopsis == 0 && game.room.backpack == 0 && game.room.key == 0 {
			result = "пустая комната. можно пройти - коридор"
		}
		if game.room.synopsis == 0 && game.room.backpack == 0 && game.room.key == 1 {
			result = "на столе: ключи. можно пройти - коридор"
		}
	}
	return result
 }

func Walk_hallway(game *game) string {
	if (*game).Kitchen.status == 1 || (*game).room.status == 1 {
		if (*game).Kitchen.status == 1 {
			(*game).Kitchen.status = 0
			(*game).hallway.status = 1
		}
		if (*game).room.status == 1 {
			(*game).room.status = 0
			(*game).hallway.status = 1
		}
		return "ничего интересного. можно пройти - кухня, комната, улица"
	} else {
		return "нет пути в коридор"
	}
 }

 func Walk_room(game *game) string {
	if (*game).hallway.status == 1{
		(*game).room.status = 1
		(*game).hallway.status = 0
		return "ты в своей комнате. можно пройти - коридор"
	} else {
		return "нет пути в комната"
	}
 }

 func Walk_kitchen(game *game) string {
	if (*game).hallway.status == 1{
		(*game).hallway.status = 0
		(*game).Kitchen.status = 1
		return "кухня, ничего интересного. можно пройти - коридор"
	}else{
		return "нет пути на кухню"
	}
 }

 func Walk_street(game *game) string {
	if (*game).hallway.status == 1{
		if (*game).inventory.key == 1{
			(*game).hallway.status = 0
			(*game).street.status = 1
			return "на улице весна. можно пройти - домой"
		} else{
			return "нет предмета в инвентаре - ключи"
		}
	} else {
		return "нет пути на улицу"
	}
 }

 func Take_key(game *game) string {
	if (*game).room.status == 1{
		if (*game).room.key == 1{
			if (*game).inventory.backpack == 1{
				(*game).room.key = 0
				(*game).inventory.key = 1
				return "предмет добавлен в инвентарь: ключи"
			}else{
				return "некуда класть"
			}
		}else{
			return "предмет уже в инвентаре"
		}
	}else{
		return "нет такого"
	}
 }

 func Put_on_backpack(game *game) string {
	if (*game).room.status == 1{
		if (*game).room.backpack == 1{
			(*game).room.backpack = 0
			(*game).inventory.backpack = 1
			return "вы надели: рюкзак"
		}else{
			return "вы уже надели рюкзак"
		}
	}else {
		return "нет такого"
	}
 }

 func Take_synopsis(game *game) string {
	if (*game).room.status == 1{
		if (*game).room.synopsis == 1{
			if (*game).inventory.backpack == 1{
				(*game).room.synopsis = 0
				(*game).inventory.synopsis = 1
				return "предмет добавлен в инвентарь: конспекты"
			}else{
				return "некуда класть"
			}
		}else{
			return "предмет уже в инвентаре"
		}
	}else{
		return "нет такого"
	}
 }