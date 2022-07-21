package main
import (
	"fmt"
)

func main()  {
	const lightSpeed = 299792
	const secondPerday = 86400
	var distance int64 = 41.3e12
	fmt.Println("alpha centauri is",distance,"km away.")
	days := distance /lightSpeed/secondPerday
	fmt.Println("that is",days,"days of travel at light speed")
}
