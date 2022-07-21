package main

import (
	"fmt"
	"strings"
)

func main()  {
	message := "your message goes here"
	keyword := "golang"
	keyindex := 0
	ciphertext := ""

	message = strings.ToUpper(strings.Replace(message," ","",-1))
	keyword = strings.ToUpper(strings.Replace(keyword," ","",-1))
	for i:=0;i<len(message);i++{
		c := message[i]
		if c >= 'A'&&c<='Z'{
			c -= 'A'
			k := keyword[keyindex]-'A'
			c = (c+k)%26 +'A'

			keyindex++
			keyindex %= len(keyword)
		}
		ciphertext += string(c)

	}
	fmt.Println(ciphertext)

}