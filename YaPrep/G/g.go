package main

import "fmt"

type Point struct {
	x, y int
}

type Path struct {
	from, to int
}

func scanCities(citiesCoordinates []*Point) {
	for i := range citiesCoordinates {
		var x, y int
		fmt.Scan(&x, &y)

		citiesCoordinates[i] = &Point{
			x: x,
			y: y,
		}
	}
}

func checkDestination()

func main() {
	var citiesNumber int

	fmt.Scan(&citiesNumber)

	citiesCoordinates := make([]*Point, citiesNumber)

	scanCities(citiesCoordinates)

	var maxPathWithoutRefuel int

	fmt.Scan(&maxPathWithoutRefuel)

	var from, to int

	fmt.Scan(&from, &to)

	path := &Path{
		from: from,
		to:   to,
	}

}
