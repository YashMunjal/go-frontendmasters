package main

import "fmt"
func average(scores [5]float64) float64 {
	total:=0.0
	for _,num :=range scores {
		total+=num
	}
	return total/float64(len(scores))
}

func main(){
	scores := [5]float64{2,7,8,9}
	fmt.Println(scores)
	fmt.Println(float64(average))
}