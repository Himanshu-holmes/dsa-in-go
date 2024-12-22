package array

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

/*
TestAddTwoNumbers tests solution(s) with the following signature and problem description:

	AddTwoNumbers(num1, num2 []int) []int

Given two numbers as an array like [2,9] and [9,9,9] return the sum of the numbers they
represent like [1,0,2,8], because 29+999=1028.
*/
func AddTwoNumbers(num1,num2 []int) []int {
	// convert num1 to string
	var num1StringArr []string;
   
	for _, num := range num1 {
		num1StringArr = append(num1StringArr, strconv.Itoa(num));
	}
	num1JoinStr := strings.Join(num1StringArr, "");
	
	
	// convert num3asstring to int
	num1AsInt, _ := strconv.Atoi(num1JoinStr);
	

	// convert num2 to string
	var  num2StringArr []string
	for _,num := range num2 {
		num2StringArr = append(num2StringArr, strconv.Itoa(num))
	};
   
	num2JoinStr := strings.Join(num2StringArr,"");

	num2AsInt,_ := strconv.Atoi(num2JoinStr);


	addBothNum := num1AsInt + num2AsInt;
	addBothNumStr := strconv.Itoa(addBothNum);
	// addBothNumStrTest := strconv.Itoa(123)
	num3Arr := strings.Split(addBothNumStr,"");


	fmt.Println("num to str",addBothNumStr)
	

	// addBothNumArr := strings.Split(addBothNumStr,"")
	fmt.Println("added num",addBothNum);

	var num3 []int;

	for _, num:= range num3Arr {
		numInt,_ := strconv.Atoi(num);
		num3 = append(num3, numInt)
	}
    
	return num3
}
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		num1, num2, sum []int
	}{
		{[]int{1}, []int{}, []int{1}},
		{[]int{1}, []int{0}, []int{1}},
		{[]int{1}, []int{1}, []int{2}},
		{[]int{1}, []int{9}, []int{1, 0}},
		{[]int{2, 5}, []int{3, 5}, []int{6, 0}},
		{[]int{2, 9}, []int{9, 9, 9}, []int{1, 0, 2, 8}},
		{[]int{9, 9, 9}, []int{9, 9, 9}, []int{1, 9, 9, 8}},
	}
	for i, test := range tests {
		if got := AddTwoNumbers(test.num1, test.num2); !slices.Equal(got, test.sum) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sum, got)
		}
	}
}
