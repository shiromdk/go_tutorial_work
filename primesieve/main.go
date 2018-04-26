package main

import (
  "fmt"
  "math"
  "time"
)



func main(){

  start := time.Now()
  var num int
  num = 100000000

  primeList := []int{}

  notPrime := make([]bool, num)
  notPrime[0] = true;
  notPrime[1] = true;
  for i:=2; i<int(math.Sqrt(float64(num)));i++{
    for j:= i+i; j<num; j=i+j{

      notPrime[j] = true
    }
  }

  for i,v := range notPrime{
    if(!v){
     primeList = append(primeList,i)
    }
  }

  fmt.Println(primeList)
  elapsed := time.Since(start)
  fmt.Println("Binomial took %s", elapsed)
}
