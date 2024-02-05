package main

import (
	"fmt"
)

type Person struct {
	FirstName   string
	LastName    string
	Age         int
	number      string
	weapons     string
	achievement string
}

func main() {
	Abl := Person{
		FirstName:   "Abylaikhan",
		LastName:    "Ybraev",
		Age:         20,
		number:      "+7 775 922 0142",
		weapons:     "AUG",
		achievement: "Топтыга Больше всего громких шагов",
	}
	Erkesh := Person{
		FirstName:   "Erkegali",
		LastName:    "Nurmukhambet",
		Age:         20,
		number:      "+7 747 847 8453",
		weapons:     "AWP",
		achievement: "Хороший стрелок! Больше попаданий в голову: [99%]",
	}
	Ili := Person{
		FirstName:   "Iliyas",
		LastName:    "Dosymbekuly",
		Age:         20,
		number:      "+7 708 149 0663",
		weapons:     "M4A4",
		achievement: "Триада!  Больше всего тройных убийств:[3]",
	}
	Dancho := Person{
		FirstName:   "Duman",
		LastName:    "Bakytkali",
		Age:         19,
		number:      "+7 747 304 3420",
		weapons:     "Desert Eagle",
		achievement: "Бережливый Победа без покупок: [3]",
	}
	Sancho := Person{
		FirstName:   "Sanjar",
		LastName:    "Serikuly",
		Age:         20,
		number:      "+7 707 309 9384",
		weapons:     "AK-47",
		achievement: "Старожил Наибольшее время жизни: [1265]",
	}

	fmt.Printf("%+v\n", Sancho)
	fmt.Printf("%+v\n", Abl)
	fmt.Printf("%+v\n", Dancho)
	fmt.Printf("%+v\n", Ili)
	fmt.Printf("%+v\n", Erkesh)

	fmt.Println("Terrorists win!")
	fmt.Println(Abl.achievement, Abl.FirstName)
	fmt.Println(Erkesh.achievement, Erkesh.FirstName)
	fmt.Println(Ili.achievement, Ili.FirstName)
	fmt.Println(Dancho.achievement, Dancho.FirstName)
	fmt.Println(Sancho.achievement, Sancho.FirstName)

}
