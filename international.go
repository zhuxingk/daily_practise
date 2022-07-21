package main

import "fmt"

func main()  {
	message := "Hola Estación Espacial Internacional"
	for _, c :=range message{
		//c := message[i]
		if c >= 'a' && c <= 'z'  {
			c = c-3
			if   c < 'a'{
				c = c-26
			}
		}else if c >= 'A' && c <= 'Z'{
			c = c-3
			if c < 'A'{
				c = c-26
			}
		}
		fmt.Printf(" %c",c)
	}
}