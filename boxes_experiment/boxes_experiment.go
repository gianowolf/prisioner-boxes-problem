package boxesexperiment_test

import (
	"fmt"

	"github.com/gianowolf/prisioner-boxes-problem/slices_manager"
)

func NewPrisionersExperiment(number_of_prisioners int) {

	boxes := slices_manager.NewUnorderedSlice(number_of_prisioners)

	for prisioner := 0; prisioner < number_of_prisioners; prisioner++ {

		fmt.Println()
		fmt.Println("Prisionero ", prisioner)
		max_attemps := size / 2
		attemp := 1

		paper := boxes[prisioner]
		fmt.Printf("Intento %d: Encontro el papel %d en la caja con su numero \n", attemp, paper)
		next_box := paper

		for {
			attemp++
			paper = boxes[next_box]
			fmt.Printf("Intento %d: Encontro el papel %d en la caja %d \n", attemp, paper, next_box)

			if prisioner == paper {
				fmt.Println("Encontro el papel con su numero en la caja", next_box)
				break
			}

			next_box = paper

			if attemp >= max_attemps {
				fmt.Println("El prisionero llego su limite de intentos")
				break
			}
		}
	}
}

func NewGroupsExperiment(prisioners_number int)
