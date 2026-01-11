package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"study/greeting"
	"time"

	"github.com/k0kubun/pp"
)

type Person struct {
	Name   string
	Age    int
	Blance int
}

func (p Person) GetName() string {
	return p.Name
}

// 2 указателя под капотом на тип объекта и сам объект
type User interface {
	GetName() string
}

func CheckUser(u User) {
	fmt.Printf(`функция которая принимает интерфейс User %v`, u.GetName(), "\n")
	fmt.Println("User Name:", u.GetName())
}

func Pay(user *Person, usd int) (int, error) {
	random := rand.IntN(100)
	pp.Println(random)
	if user.Blance-usd < 0 {
		return 0, errors.New("Недостаточно средств!")
	}

	return random, nil
}

func main() {

	defer func() {
		panic := recover()

		if panic != nil {
			fmt.Println("Произошла ошибка:", panic)
		}
	}()

	var person = Person{
		Name:   "Alice",
		Age:    30,
		Blance: 300,
	}
	val, err := Pay(&person, 100)
	if err != nil {
		fmt.Println("Ошибка при оплате:", err)
	} else {
		fmt.Println("Оплата прошла успешно, ваш код:", val)
	}

	fmt.Println("Hello, 世界")

	greeting.SayGf()

	checkList := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 7,
	}

	value, err2 := checkList["apple2"]

	if err2 != true {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Success")
	}

	fmt.Println(value, err2)

	// any
	//var name string = "mazaka"
	//const nix = 66.7
	//fmt.Printf("%T\n", nix)

	whatAmI := func(item any) {
		switch t := item.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI("hey")

	intChan := make(chan int)

	go func() {
		intChan <- 666
	}()

	select {
	case msg := <-intChan:
		fmt.Println(msg)
	}

	mapa := map[string]int{
		"dd": 21,
		"ff": 22,
	}
	mapa["x"] = 2
	val, ok := mapa["yy"]
	if ok != true {
		fmt.Println("Error:", ok)
		pp.Println("mapa", true)
	}

	var ch = make(chan string)

	go func() {
		ch <- "hello"
	}()
	var vvv = <-ch
	fmt.Println(vvv)

	type Message struct {
		Name string
	}

	var robot = []any{}

	// создаю новый открытый канал
	messageChan := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		for i := 0; i < 6; i++ {
			messageChan <- "name" + strconv.Itoa(i)
		}
		// закрываем канал
		close(messageChan)
	}()

	// // блокируем поток и ждем значение от канала
	// v1, ok := <-messageChan
	// fmt.Println(v1, ok)

	for mess := range messageChan {
		robot = append(robot, mess)
	}

	pp.Println("robot", robot)

}
