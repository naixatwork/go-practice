package main

import "fmt"

const (
	maxX = 1000
	maxY = 600
)

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("%#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		X: 10,
		Y: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(newItem(10, 20))
	fmt.Println(newItem(10, 2-0))

	i3.Move(100, 200)
	fmt.Printf("i3: %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: i3,
	}
	p1.Move(400, 600)
	i3.Move(100, 140)
	fmt.Println(p1)
	fmt.Println(i3)
}

type Player struct {
	Name string
	Item
}

type Item struct {
	X int
	Y int
}

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

func newItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}

	return &i, nil
}
