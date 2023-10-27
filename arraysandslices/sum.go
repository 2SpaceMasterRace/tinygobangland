package main

import "fmt"

func Sum(numbers []int)int {
  sum := 0
  for _, number:= range numbers{
    sum += number
  }
  return sum
}

func SumAll(numbersToSum ...[]int)[]int{
  var sums []int
  for _,numbers:= range numbersToSum{
    sums = append(sums, Sum(numbers))
  }
  return sums
}

func SumTails(numbersToSum ...[]int) []int  {
  var sumtails []int
  for _,numbers := range numbersToSum{
    if len(numbers) == 0{
      sumtails = append(sumtails, 0)
    }else {
      tails := numbers[1:]
      sumtails = append(sumtails, Sum(tails))
    }
   }
  return sumtails
}

func main()  {
 fmt.Println(SumAll([]int{0,5},[]int{1,3}))
}
