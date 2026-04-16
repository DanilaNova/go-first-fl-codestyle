package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func doAttack(charName, charClass string) string {
	var attack_power int
	switch charClass {
	case "warrior":
		attack_power = 5 + randInt(3, 5)
	case "mage":
		attack_power = 5 + randInt(5, 10)
	case "healer":
		attack_power = 5 + randInt(-3, -1)
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, attack_power)
}

func doDefence(char_name, charClass string) string {
	var defence_power int
	switch charClass {
	case "warrior":
		defence_power = 10 + randInt(5, 10)
	case "mage":
		defence_power = 10 + randInt(-2, 2)
	case "healer":
		defence_power = 10 + randInt(2, 5)
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s блокировал %d урона.", char_name, defence_power)
}

func doSpecial(charName, charClass string) string {
	var special string
	switch charClass {
	case "warrior":
		special = fmt.Sprintf("Выносливость %d", 80+25)
	case "mage":
		special = fmt.Sprintf("Атака %d", 5+40)
	case "healer":
		special = fmt.Sprintf("Защита %d", 10+30)
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s применил специальное умение `%s`", charName, special)
}

func startTraining(charName, charClass string) string {

	switch charClass {
	case "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		switch cmd {
		case "attack":
			fmt.Println(doAttack(charName, charClass))
		case "defence":
			fmt.Println(doDefence(charName, charClass))
		case "special":
			fmt.Println(doSpecial(charName, charClass))
		}
	}

	return "тренировка окончена"
}

func chooseCharClass() string {
	var approveChoice string
	var class string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &class)

		switch class {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			fmt.Println("Нейзвестный класс")
			continue
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return class
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := chooseCharClass()

	fmt.Println(startTraining(charName, charClass))
}
