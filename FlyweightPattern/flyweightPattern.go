package main

import "fmt"

type Character interface {
	Draw(int, int)
}

type ConcreteCharacter struct {
	symbol string
	font   string
	size   int
}

func (c *ConcreteCharacter) Draw(x int, y int) {
	fmt.Printf("Drawing %s of style %s and size %d at (%d, %d)\n", c.symbol, c.font, c.size, x, y)
}

type CharacterFactory struct {
	ch map[string]*ConcreteCharacter
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		ch: make(map[string]*ConcreteCharacter),
	}
}

func (c *CharacterFactory) GetCharacter(symbol string) *ConcreteCharacter {
	if ch, ok := c.ch[symbol]; ok {
		fmt.Printf("%s already present in the cache\n", symbol)
		return ch
	}

	ch := &ConcreteCharacter{
		symbol: symbol,
		font:   "Times New Roman",
		size:   10,
	}
	// add the newly created concrete character to character factory
	c.ch[symbol] = ch
	fmt.Printf("%s is added to the cache\n", symbol)
	return ch
}

func main() {
	factory := NewCharacterFactory()
	concreteChar1 := factory.GetCharacter("A")
	concreteChar2 := factory.GetCharacter("A")
	concreteChar3 := factory.GetCharacter("B")
	concreteChar1.Draw(1, 2)
	concreteChar2.Draw(2, 3)
	concreteChar3.Draw(4, 10)
}
