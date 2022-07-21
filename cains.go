package main
import "fmt"
import "math/big"

func main()  {
	lightspeed := big.NewInt(299792)
	secondsPeryear := big.NewInt(60*60*24*365)
	distance :=	new(big.Int)
	distance.SetString("236000000000000000000",10)
	lightsecondsPeryear := new(big.Int)
	lightsecondsPeryear.Div(distance,lightspeed)
	lightyear := new(big.Int)
	lightyear.Div(lightsecondsPeryear,secondsPeryear)
	fmt.Printf("this is distance is %s lightyear!",lightyear)
}