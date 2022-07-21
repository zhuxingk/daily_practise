package main

import "fmt"

//func main()  {
//	c := 'a'
//	c = c+6
//	fmt.Printf("%c\n",c)
//	c = c-'a'+'A'
//	fmt.Printf("%c\n",c)
//}
func main()  {
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0;i<len(message);i++{
		c := message[i]
		if c >= 'a' && c <= 'z'{
			c = c +13
			if c > 'z'{
				c = c-26
			}
			}
		fmt.Printf("%c",c)
	}

}