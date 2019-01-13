package main

import (
	"fmt"
)

// インターフェース定義
type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Println("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"picomaru"}
	talker.Talk()
}
