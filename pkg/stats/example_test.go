package stats

import (
	"github.com/s-zer0/bank/v2/pkg/types"
	"fmt")

func ExampleAvg(){
	payments := []types.Payment{
	{
	    ID:2,
      Amount:53_00,
      Category: "Cat",
      Status: "FAIL",

	},
	
    {
      ID:1,
      Amount:51_00,
      Category: "Cat",
      Status: "OK",
	},
	
    {
      ID:3,
      Amount:52_00,
      Category: "Cat",
      Status: "INPROGRESS",
    },
  }

  fmt.Println(Avg(payments))

  //Output: 5150
}

func ExampleTotalInCategory(){
  payments := []types.Payment{
    {
      ID:2,
      Amount:53_00,
      Category: "Cafe",
      Status: "INPROGRESS",
    },
    {
      ID:1,
      Amount:51_00,
      Category: "Cafe",
      Status: "OK",
    },
    {
      ID:3,
      Amount:52_00,
      Category: "Restaurant",
      Status: "FAIL",
    },
  }

  fmt.Println(TotalInCategory(payments, "Cafe"))

  //Output: 10400
}