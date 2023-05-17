package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type EaterFood struct {
	eater_id    string
	foodmenu_id string
}

type FoodCount struct {
	foodmenu_id string
	count       int
}

func convert_to_struct_data_and_count(logFile string) []FoodCount {
	data, err := ioutil.ReadFile(logFile)
	if err != nil {
		log.Fatal(err)
	}
	is_there := make(map[string]bool)
	count_map := make(map[string]int)

	var all_data []EaterFood

	for _, x := range strings.Split(string(data), "\n") {
		if x != "" {
			// fmt.Println(x)
			var data EaterFood
			ids := strings.Split(x, ",")

			data.eater_id = ids[0]
			data.foodmenu_id = ids[1]
			key := fmt.Sprintf("%s:%s", data.eater_id, data.foodmenu_id)
			if !is_there[key] {
				all_data = append(all_data, data)
				count_map[data.foodmenu_id] += 1
				is_there[key] = true
			} else {
				fmt.Fprintln(os.Stderr, "this eater", data.eater_id, "has duplicate with food id =", data.foodmenu_id)
			}

			// fmt.Printf("eater %s, food %s\n", data.eater_id, data.foodmenu_id)

		}

	}
	// fmt.Printf("%#v\n", count_map)
	food_counts := []FoodCount{}
	for foodmenu_id, count := range count_map {
		var food_count FoodCount
		food_count.count = count
		food_count.foodmenu_id = foodmenu_id

		food_counts = append(food_counts, food_count)
	}
	return food_counts

}

func bubble_sort(food_counts []FoodCount) []FoodCount {
	for i := 0; i < len(food_counts)-1; i++ {
		for j := 0; j < len(food_counts)-i-1; j++ {
			if food_counts[j].count < food_counts[j+1].count {
				food_counts[j], food_counts[j+1] = food_counts[j+1], food_counts[j]
			}
		}
	}
	return food_counts

}

func top_three(food_counts []FoodCount) []FoodCount {

	top3 := make([]FoodCount, 0, 3)

	for i, x := range food_counts {
		if i < 3 {
			top3 = append(top3, x)
		} else {
			break
		}
	}
	return top3
}

func main() {
	logFile := "log.txt"

	food_counts := convert_to_struct_data_and_count(logFile)

	food_counts = bubble_sort(food_counts)
	top_3 := top_three(food_counts)
	fmt.Println("Top 3 food")
	for _, x := range top_3 {
		fmt.Println(x.foodmenu_id)
	}

}
