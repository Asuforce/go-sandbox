package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type placeholder [5]string

var (
	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	colon = placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}
	digits = [...]placeholder{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}
)

func main() {
	screen.Clear()

	for {
		screen.MoveTopLeft()

		n := time.Now()
		h, m, s := n.Hour(), n.Minute(), n.Second()

		clock := [...]placeholder{
			digits[h/10], digits[h%10],
			colon,
			digits[m/10], digits[m%10],
			colon,
			digits[s/10], digits[s%10],
		}

		for l := range clock[0] {
			for i, v := range clock {
				next := clock[i][l]
				if v == colon && s%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
