package main

import "fmt"

type CElsius float64

func (c CElsius) fahrenheit() fahrenheit {//celsius to fahrenheit
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c CElsius) KElvin() KElvin {//celsius to kelvin
	return KElvin(c + 273.15)
}

type fahrenheit float64

func (f fahrenheit) CElsius() CElsius {//fahrenheit to celsius
	return CElsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) KElvin() KElvin {//fahrenheit to kelvin
	return f.CElsius().KElvin()
}

type KElvin float64

func (k KElvin) CElsius() CElsius {//kelvin to celsius
	return CElsius(k - 273.15)
}

func (k KElvin) fahrenheit() fahrenheit {//kelvin to fahrenheit
	return k.CElsius().fahrenheit()
}

func main() {
	var k KElvin = 294.0
	c := k.CElsius()
	fmt.Print(k, "ºK is ", c, "ºC")
}

