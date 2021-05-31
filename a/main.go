package a

import (
	"github.com/enmand/synk-fixtures-go-mvs/b"
	"github.com/enmand/synk-fixtures-go-mvs/c"
	"github.com/enmand/synk-fixtures-go-mvs/e"
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
