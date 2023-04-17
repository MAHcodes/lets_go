package main

import "fmt"

func sum(nums []int, c chan int)  {
  s := 0
  for _, v := range nums {
    s += v
  }
  c <- s
}

func main () {
  c := make(chan int)
  nums := []int{1,2,4,5,6,8}

  go sum(nums[:len(nums)/2], c)
  go sum(nums[len(nums)/2:], c)

  x := <- c
  y := <- c
  go sum([]int{x, y}, c)

  z := <- c
  fmt.Println(z)
}
