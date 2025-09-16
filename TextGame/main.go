package main

import (
	"strings"
)

type Room struct {
	Table []string
	Chair []string
}

var current string
var rooms map[string]*Room
var inventory map[string]bool
var rucksackWorn bool
var doorOpen bool

func initGame() {
	current = "кухня"
	rooms = map[string]*Room{
		"кухня":   {Table: []string{"чай"}, Chair: []string{}},
		"комната": {Table: []string{"ключи", "конспекты"}, Chair: []string{"рюкзак"}},
		"коридор": {Table: []string{}, Chair: []string{}},
		"улица":   {Table: []string{}, Chair: []string{}},
	}
	inventory = map[string]bool{}
	rucksackWorn = false
	doorOpen = false
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func remove(slice []string, s string) []string {
	out := []string{}
	for _, v := range slice {
		if v != s {
			out = append(out, v)
		}
	}
	return out
}

func describeCurrent() string {
	switch current {
	case "кухня":
		if !rucksackWorn {
			return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		}
		return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
	case "комната":
		t := rooms["комната"].Table
		c := rooms["комната"].Chair
		if len(t) == 0 && len(c) == 0 {
			return "пустая комната. можно пройти - коридор"
		}
		parts := []string{}
		if len(t) > 0 {
			parts = append(parts, "на столе: "+strings.Join(t, ", "))
		}
		if len(c) > 0 {
			parts = append(parts, "на стуле: "+strings.Join(c, ", "))
		}
		return strings.Join(parts, ", ") + ". можно пройти - коридор"
	case "коридор":
		return "ничего интересного. можно пройти - кухня, комната, улица"
	case "улица":
		return "на улице весна. можно пройти - домой"
	}
	return ""
}

func handleCommand(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return "неизвестная команда"
	}
	switch parts[0] {
	case "осмотреться":
		return describeCurrent()
	case "идти":
		if len(parts) < 2 {
			return "неизвестная команда"
		}
		target := parts[1]
		switch target {
		case "коридор":
			// из кухни и из комнаты можно идти в коридор
			if current == "кухня" || current == "комната" || current == "коридор" {
				current = "коридор"
				return "ничего интересного. можно пройти - кухня, комната, улица"
			}
			return "нет пути в коридор"
		case "комната":
			if current == "коридор" {
				current = "комната"
				return "ты в своей комнате. можно пройти - коридор"
			}
			return "нет пути в комната"
		case "кухня":
			if current == "коридор" {
				current = "кухня"
				return "кухня, ничего интересного. можно пройти - коридор"
			}
			return "нет пути в кухня"
		case "улица":
			if current == "коридор" {
				if doorOpen {
					current = "улица"
					return "на улице весна. можно пройти - домой"
				}
				return "дверь закрыта"
			}
			return "нет пути в улица"
		}
	case "надеть":
		if len(parts) < 2 {
			return "неизвестная команда"
		}
		if parts[1] == "рюкзак" {
			if contains(rooms[current].Chair, "рюкзак") {
				rooms[current].Chair = remove(rooms[current].Chair, "рюкзак")
				rucksackWorn = true
				return "вы надели: рюкзак"
			}
			// если он уже надет — просто вернуть ту же строку
			if rucksackWorn {
				return "вы надели: рюкзак"
			}
			return "нет такого"
		}
	case "взять":
		if len(parts) < 2 {
			return "неизвестная команда"
		}
		item := parts[1]
		// попытка взять предмет со стола
		if contains(rooms[current].Table, item) {
			if !rucksackWorn {
				return "некуда класть"
			}
			inventory[item] = true
			rooms[current].Table = remove(rooms[current].Table, item)
			return "предмет добавлен в инвентарь: " + item
		}
		// попытка взять со стула
		if contains(rooms[current].Chair, item) {
			// специфики: если это рюкзак — лучше надеть, но мы позволим взять
			if !rucksackWorn && item != "рюкзак" {
				return "некуда класть"
			}
			inventory[item] = true
			rooms[current].Chair = remove(rooms[current].Chair, item)
			return "предмет добавлен в инвентарь: " + item
		}
		return "нет такого"
	case "применить":
		if len(parts) < 3 {
			return "неизвестная команда"
		}
		item := parts[1]
		target := parts[2]
		if !inventory[item] {
			return "нет предмета в инвентаре - " + item
		}
		if item == "ключи" && target == "дверь" {
			doorOpen = true
			return "дверь открыта"
		}
		return "не к чему применить"
	}
	return "неизвестная команда"
}

func main() {
}
