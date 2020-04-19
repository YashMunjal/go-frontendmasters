package main

import ( "fmt"
)

//part 1
func average(scores [5]float64) float64 {
	total:=0.0
	for _,num :=range scores {
		total+=num
	}
	return total/float64(len(scores))
}

//part 2

var initialPets map[string]string= map[string]string{
	"saksham":"mosquito",
	"bansal":"panda",
	"deepankur":"scorpion",
	"vaibhavMittal":"doggo",

}
func doesPetExists(petName string) bool{
	_,ok :=initialPets[petName]
	return ok
}
func main(){
	scores := [5]float64{2,7,8,9}
	//fmt.Println(scores)
	fmt.Println(average(scores))

	fmt.Println(doesPetExists("bansal"))
	//returns true
}



