package main

import (
	"fmt"
)

func test1() {

	k1 := 1.0
	d1 := 0.0

	k2 := -1.0
	d2 := +2.0

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	}
}

func test2() {

	k1 := 3.0 / 2.0
	d1 := 0.0

	k2 := -1.0
	d2 := +2.0

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	}
}

func test3() {

	k1 := 2.0
	d1 := 1.0

	k2 := +2.0
	d2 := -1.0

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no section - lines parallel")
	}
}

func test4() {

	vertices := createCircle(0.5, 0.5, 0.5, 8)
	x, y, ok := tryCalcMiddlePoint(vertices, 1)

	if ok {
		fmt.Printf("middlepoint x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no middlepoint")
	}
}

func test5() {

	vertices := []float64{
		2.0, 7.0,
		3.0, 4.0,
		8.0, 3.0}

	x, y, ok := tryCalcMiddlePoint(vertices, 1)

	if ok {
		fmt.Printf("middlepoint x: %f y: %f", x, y)
		fmt.Println()
		r, ok := tryCalcRadius(vertices, 1)
		if ok {
			fmt.Printf("radius: %f", r)
			fmt.Println()
			fmt.Printf("kruemmung: %f at x: 2.0, y: 7.0", 1/r)
			fmt.Println()
			v := 50.0  // m/s = 180 km/h
			m := 700.0 // kg
			fg := m * v * v * (1.0 / float64(r))
			fmt.Printf("Fg: %f kN", fg/1000.0)
			fmt.Println()
		}

	} else {
		fmt.Println("no middlepoint")
	}
}

func test6() {

	x1 := 0.0
	y1 := 0.0
	x2 := 0.0
	y2 := 1.0

	k, d := calcLineSymetrale(x1, y1, x2, y2)

	fmt.Printf("k: %f, d: %f \n", k, d)
}

func test7() {

	x1 := 0.5
	y1 := 2.5
	x2 := 1.5
	y2 := 0.5

	k, d := calcLineSymetrale(x1, y1, x2, y2)

	fmt.Printf("k: %f, d: %f \n", k, d)
}

func tests() {

	/*
		test1()
		test2()
		test3()
		test4()
		test5()*/
	//test6()
	test7()
}
