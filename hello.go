package main

import (
	"fmt"
	"math"
)

func sign(x1, y1, x2, y2, x3, y3 int) float64 {
	return float64(x1-x3)*float64(y2-y3) - float64(x2-x3)*float64(y1-y3)
}

func pointInTriangle(X, Y [3]int, x, y int) bool {
	var d1, d2, d3 float64
	var hasNeg, hasPos bool

	d1 = sign(x, y, X[0], Y[0], X[1], Y[1])
	d2 = sign(x, y, X[1], Y[1], X[2], Y[2])
	d3 = sign(x, y, X[2], Y[2], X[0], Y[0])

	hasNeg = (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos = (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(hasNeg && hasPos)
}

func pointsAreCollinear(x1, y1, x2, y2, x3, y3 int) bool {
	return (y2-x2)*(x3-y1) == (y3-y2)*(x2-x1)
}

//func distance(x1, y1)

func pointIsInTheMiddle(x1, y1, x2, y2, x, y int) bool {
	//((x1 <= x && x <= x2) || (x2 <= x && x <= x1)) && ((y1 <= y && y <= y2) || (y2 <= y && y <= y1)) OLD VERSION NOT WORKING
	//fmt.Printf("Points 2 is in the middle of 0 and 1?: x[0] = %v y[0] = %v, x[1] = %v y[1] = %v, x[2] = %v y[2] = %v. \n", x1, y1, x2, y2, x, y)
	return (x1 == x && x == x2 && (y1 < y && y < y2) || (y1 > y && y > y2)) ||
		(y1 == y && y == y2 && (x1 < x && x < x2) || (x1 > x && x > x2)) ||
		(x1 < x && x < x2 && y1 < y && y < y2) ||
		(x1 > x && x > x2 && y1 > x && y > y2) ||
		(x1 < x && x < x2 && y1 > y && y > y2) ||
		(x1 > x && x > x2 && y1 < y && y < y2)
}

func areaOfTriangle(x1, y1, x2, y2, x3, y3 int) float64 {
	return 0.5 * (math.Abs(float64(x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))))
}

func pointsContainPoint(X, Y [3]int, x, y int) bool {
	for i := 0; i < len(X); i++ {
		if X[i] == x && Y[i] == y {
			return true
		}
	}
	return false
}

func main() {
	X := [12]int{0, 0, 0, 5, 4, 8, 1, 4, 3, -4, -2, 2}
	Y := [12]int{0, 2, 4, -3, 5, 7, -3, 0, 5, -4, 2, 13}

	fmt.Println(X)
	fmt.Println(Y)

	x := [3]int{X[0], X[1]}
	y := [3]int{Y[0], Y[1]}

	//var x1, x2, x3 = X[0], X[1], X[2]
	//var y1, y2, y3 = Y[0], Y[1], Y[2]

	for j := 2; j < len(X); j++ {
		x[2] = X[j]
		y[2] = Y[j]
		//fmt.Printf("Points are Collinear?: x[0] = %v y[0] = %v, x[1] = %v y[1] = %v, x[2] = %v y[2] = %v \n", x[0], y[0], x[1], y[1], x[2], y[2])
		if pointsAreCollinear(x[0], y[0], x[1], y[1], x[2], y[2]) {
			//var k int;
			//fmt.Printf("Points are Collinear: x[0] = %v y[0] = %v, x[1] = %v y[1] = %v, x[2] = %v y[2] = %v \n", x[0], y[0], x[1], y[1], x[2], y[2])
			if pointIsInTheMiddle(x[0], y[0], x[1], y[1], x[2], y[2]) {
				//	k = 2
				//	fmt.Printf("Points 2 is in the middle of 0 and 1: x[0] = %v y[0] = %v, x[1] = %v y[1] = %v, x[2] = %v y[2] = %v \n", x[0], y[0], x[1], y[1], x[2], y[2])
				x[1] = x[2]
				y[1] = y[2]
			}
		} else {
			break
		}
	}

	for i := 0; i < len(X); i++ {
		if pointsContainPoint(x, y, X[i], Y[i]) {
			continue
		} else if pointInTriangle(x, y, X[i], Y[i]) {
			//fmt.Printf("Point %v,%v is in triangle of points [%v,%v,%v][%v,%v,%v]\n", X[i], Y[i], x[0], x[1], x[2], y[0], y[1], y[2])
			//fmt.Println()
			if pointsAreCollinear(x[0], y[0], x[1], y[1], X[i], Y[i]) && pointIsInTheMiddle(x[0], y[0], x[1], y[1], X[i], Y[i]) {
				x[1] = X[i]
				y[1] = Y[i]
			} else /*if pointsAreCollinear(x[0], y[0], x[2], y[2], X[i], Y[i])*/ {
				x[2] = X[i]
				y[2] = Y[i]
			} /*else if pointsAreCollinear(x[1], y[1], x[2], y[2], X[i], Y[i]) {
				x[2] = X[i]
				y[2] = Y[i]
			} else {
				x[2] = X[i]
				y[2] = X[i]
			}*/
		}
	}

	if pointsAreCollinear(x[0], y[0], x[1], y[1], x[2], y[2]) {
		fmt.Println("There are no points that make a triangle with positive area and contain no other points")
	} else {
		fmt.Println(x)
		fmt.Println(y)
	}

	//fmt.Println("Hello, world.")
}
