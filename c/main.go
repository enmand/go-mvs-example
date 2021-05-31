package c

import "github.com/enmand/synk-fixtures-go-mvs/d"

func GetC() string {
	return "c:1.0.0"
}

func GetD() string {
	return d.GetD()
}
