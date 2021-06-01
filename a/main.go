package a

import (
	"github.com/enmand/go-mvs-example/b"
	"github.com/enmand/go-mvs-example/c"
	"github.com/enmand/go-mvs-example/e"
)

func GetA() string {
	return "a:1.0.0"
}

func GetB() string {
	return b.GetB()
}

func GetC() string {
	return c.GetC()
}

func GetE() string {
	return e.GetE()
}
