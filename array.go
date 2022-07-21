package main

import "fmt"

func main() {
	var team = [3]string{"hammer","soldier","mum"}
	var team1 = [...]string{"fire","gun","aex","fight"}
	fmt.Println(team)
	fmt.Println(team1)

for k, v := range team {
	fmt.Println(k, v)
}
}