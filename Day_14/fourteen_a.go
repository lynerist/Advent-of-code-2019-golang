package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name     string
	quantity int
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	reactions := make(map[string][]ingredient)
	// Here We scan the input (you can do go run program.go < input.txt)
	for {
		sc.Scan()
		if sc.Text() == "" {
			break
		}

		inputList := strings.FieldsFunc(sc.Text(), func(r rune) bool {
			if r == ' ' || r == ',' {
				return true
			}
			return false
		})
		//We add all in a map that link product -> list of things that I need to make it
		var ingredients []ingredient
		for i := 1; ; i += 2 {
			quantity, _ := strconv.Atoi(inputList[i-1])
			ingredients = append(ingredients, ingredient{inputList[i], quantity})
			if inputList[i+1] == "=>" {
				quantityResult, _ := strconv.Atoi(inputList[i+2])
				//The first element will give to us the number of result producted
				ingredients = append([]ingredient{ingredient{"", quantityResult}}, ingredients...)
				reactions[inputList[i+3]] = ingredients
				break
			}
		}
	}
	//Here we store the excess to use it when we need
	refuse := make(map[string]int)

	fmt.Println(cost("FUEL", 1, reactions, refuse))
}

func cost(name string, quantity int, reactions map[string][]ingredient, refuse map[string]int) int {
	//End of the chain
	if name == "ORE" {
		return quantity
	}

	//We separe the first element that isn't an ingredient
	list, quantityProducted := reactions[name][1:], reactions[name][0].quantity

	//We see if we have producted that thing in excess yet
	if refuse[name] > 0 {
		if refuse[name] >= quantity {
			refuse[name] -= quantity
			return 0
		}
		quantity -= refuse[name]
		refuse[name] = 0
		return cost(name, quantity, reactions, refuse)
	}

	//We round up the number of reactions needed
	reactionsNeeded := (quantity-1)/quantityProducted + 1

	//We sum the ore needed for each ingredient
	var oreNeeded int
	for _, r := range list {
		oreNeeded += cost(r.name, r.quantity*reactionsNeeded, reactions, refuse)
	}

	//Here we put the excess in the refuse map
	if reactionsNeeded*quantityProducted-quantity > 0 {
		refuse[name] += reactionsNeeded*quantityProducted - quantity
	}
	return oreNeeded
}
