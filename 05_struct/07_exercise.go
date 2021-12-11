package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
	luxury    bool
}

type sedan struct {
	vehicle
	fourWheel bool
	luxury    bool
}

func main() {

	//
	// truck
	lumberJack := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
		fourWheel: true,
		luxury:    false,
	}

	fmt.Println(lumberJack)
	fmt.Println(lumberJack.vehicle.color)
	fmt.Println(lumberJack.vehicle.doors)
	fmt.Println(lumberJack.fourWheel)
	fmt.Println(lumberJack.luxury)

	//
	// sedan
	ford := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		fourWheel: false,
		luxury:    true,
	}
	fmt.Println(ford)
	fmt.Println(ford.vehicle.color)
	fmt.Println(ford.vehicle.doors)
	fmt.Println(ford.fourWheel)
	fmt.Println(ford.luxury)

}
