package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
  t.Run("collection of 5 numbers",func(t *testing.T){
  numbers := []int{1,2,3,4,5}

  got := Sum(numbers)
  want := 15

  if got!=want{
    t.Errorf("got %d want %d, given %v", got,want,numbers)
  }
  })
  t.Run("collection of any numbers", func(t *testing.T){
    numbers := []int{1,2,3}

    got := Sum(numbers)
    want:= 6

    if got!=want{
      t.Errorf("got %d want %d given %v", got,want,numbers)
    }
  })
}

func TestSumAll(t *testing.T){
    got := SumAll([]int{1,2}, []int{0,9})
    want := []int{3,9}

    if !reflect.DeepEqual(got, want){
      t.Errorf("got %v want %v", got, want)
    }
}

func TestSumTails(t *testing.T){
  t.Run("check: valid inputs", func(t *testing.T){
    got:= SumTails([]int{1,9,1}, []int{2,4,1})
    want := []int{10,5}

    if !reflect.DeepEqual(got, want){
      t.Errorf("got %v want %v", got,want)
  }
  })
  t.Run("check: empty slice", func(t *testing.T){
    got := SumTails([]int{},[]int{1,2,3})
    want := []int{0,5}

    if !reflect.DeepEqual(got, want){
      t.Errorf("got %v want %v", got,want)
    }
  })
}
