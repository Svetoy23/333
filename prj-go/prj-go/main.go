package main

import (
	domain "command-line-argumentsC:\\Users\\NUPP\\Documents\\VE201\\prj-go\\prj-go\\domain\\user.go"
	"fmt"
	"math/rand"
	"time"
)

var want string
var resz, points int

const pointstowin int = 50

var pointsperque int = 10
var id uint64 = 1

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello user! Do you want to play a game?(y/n)")
	fmt.Scan(&want)

	if want == "y" {
		load := time.NewTimer(3 * time.Second)
		fmt.Println("Loading game... 99%")
		<-load.C

		fmt.Println("The game has begun!")
		start := time.Now()
	} else {
		return
	}
}

func play() domain.User {
	for pointstowin >= points {
		a := rand.Intn(10)
		b := rand.Intn(10)
		i := rand.Intn(3) + 1

		switch i {
		case 1:
			fmt.Printf("%d + %d =\n", a, b)
			fmt.Scan(&resz)
			if a+b == resz {
				points += pointsperque
				fmt.Println("Welldone!")
			} else {
				fmt.Println("Wrong.. :(")
			}
		case 2:
			fmt.Printf("%d - %d =\n", a, b)
			fmt.Scan(&resz)
			if a-b == resz {
				points += pointsperque
				fmt.Println("Welldone!")
			} else {
				fmt.Println("Wrong.. :(")
			}
		case 3:
			fmt.Printf("%d * %d =\n", a, b)
			fmt.Scan(&resz)
			if a*b == resz {
				points += pointsperque
				fmt.Println("Welldone!")
			} else {
				fmt.Println("Wrong.. :(")
			}
		case 4:
			if b == 0 {
				fmt.Println("Oops, we tried to divide by zero. Let's try again.")
				break
			} else {
				fmt.Printf("%d / %d =\n", a, b)
				fmt.Scan(&resz)
				if a/b == resz {
					points += pointsperque
					fmt.Println("Welldone!")
				} else {
					fmt.Println("Wrong.. :(")
				}
			}
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("Час гри який був затрачений на її проходження %s\n", elapsed)
	fmt.Println("Введіть своє ім'я: ")

	name := ""
	fmt.Scan(&name)

	var user domain.User
	user.Id = id
	user.Name = name
	user.Time = start
	id++
}
