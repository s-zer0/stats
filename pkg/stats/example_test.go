package stats

import (
	"github.com/s-zer0/bank/v2/pkg/types"
	//"fmt"
	"testing"
	"reflect")

// func ExampleAvg(){
// 	payments := []types.Payment{
// 	{
// 	    ID:2,
//       Amount:53_00,
//       Category: "Cat",
//       Status: "FAIL",

// 	},
	
//     {
//       ID:1,
//       Amount:51_00,
//       Category: "Cat",
//       Status: "OK",
// 	},
	
//     {
//       ID:3,
//       Amount:52_00,
//       Category: "Cat",
//       Status: "INPROGRESS",
//     },
//   }

//   fmt.Println(Avg(payments))

//   //Output: 5150
// }

// func ExampleTotalInCategory(){
//   payments := []types.Payment{
//     {
//       ID:2,
//       Amount:53_00,
//       Category: "Cafe",
//       Status: "INPROGRESS",
//     },
//     {
//       ID:1,
//       Amount:51_00,
//       Category: "Cafe",
//       Status: "OK",
//     },
//     {
//       ID:3,
//       Amount:52_00,
//       Category: "Restaurant",
//       Status: "FAIL",
//     },
//   }

//   fmt.Println(TotalInCategory(payments, "Cafe"))

//   //Output: 10400
// }

func TestCategoriesAvgUser(t *testing.T) {
	payments := []types.Payment{
	  {ID:1, Category: "auto", Amount: 1_000_00},
	  {ID:2, Category: "food", Amount: 2_000_00},
	  {ID:3, Category: "auto", Amount: 4_000_00},
	  {ID:4, Category: "auto", Amount: 4_000_00},
	  {ID:5, Category: "fun", Amount: 5_000_00},
	} 
	expected := map[types.Category]types.Money{
	  "auto": 3_000_00,
	  "food": 2_000_00,
	  "fun": 5_000_00,
	}
  
	result := CategoriesAvg(payments)
  
	if !reflect.DeepEqual (expected, result){
	  t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
  }

  func TestPeriodsDynamicUser(t *testing.T) {
	first := map[types.Category]types.Money {
		"auto":	1_000_000,
		"food":	2_000_000,
		"fun":	5_000_000,
	}
	second := map[types.Category]types.Money {
		"auto":	2_000_000,
		"food":	3_000_000,
		"fun":	3_000_000,
		"mobile": 1_000_000,
	}

	expected := map[types.Category]types.Money{
		"auto":	1_000_000,
		"food":	1_000_000,
		"fun":	-2_000_000,
		"mobile": 1_000_000,
		
	}

	result := PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected,result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}