package main

import "testing"

func TestMenuAnalyzer(t *testing.T) {
	logFile := "log.txt"

	expectedTop3 := []FoodCount{
		{foodmenu_id: "106", count: 10},
		{foodmenu_id: "101", count: 4},
		{foodmenu_id: "102", count: 3},
	}

	foodCounts := convert_to_struct_data_and_count(logFile)
	sortedFoodCounts := bubble_sort(foodCounts)
	top3 := top_three(sortedFoodCounts)

	if len(top3) != len(expectedTop3) {
		t.Errorf("Expected top 3 count: %d, Got: %d", len(expectedTop3), len(top3))
	}

	for i := 0; i < len(top3); i++ {
		if top3[i].foodmenu_id != expectedTop3[i].foodmenu_id || top3[i].count != expectedTop3[i].count {
			t.Errorf("Expected: (%s, %d), Got: (%s, %d)", expectedTop3[i].foodmenu_id, expectedTop3[i].count, top3[i].foodmenu_id, top3[i].count)
		}
	}
}
