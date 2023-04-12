package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Item struct {
	tier string
	item string
}

func main() {
	fmt.Println("Enter the path to the text file of items you want to rank")

	pathReader := bufio.NewReader(os.Stdin) 

	path, _ := pathReader.ReadString('\n')
	path = strings.Replace(path, "\n", "", -1)

	file, err := os.Open(path) 

	if err != nil {
		log.Fatal("There was an error reading the file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var items []string

	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var tiers []Item

	for _, item := range items {
		fmt.Println(
			fmt.Sprintf(
				"What tier do you want to place %s in? (S, A, B, C, D, F)",
				item,
			),
		)

		tierReader := bufio.NewReader(os.Stdin)	

		tier, _ := tierReader.ReadString('\n')
		tier = strings.Replace(tier, "\n", "", -1)

		tiers = append(tiers, Item{
			tier: tier,
			item: item,
		})	
	}

	for _, tier := range []string{"S", "A", "B", "C", "D", "F"} {
		fmt.Println(
			fmt.Sprintf(
				"%s\n---------",
				tier,
			),
		)
		
		for _, item := range tiers {
			if strings.ToUpper(item.tier) == tier {
				fmt.Println(item.item)
			}
		}
	}
}
