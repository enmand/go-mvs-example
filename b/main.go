package b

import "github.com/enmand/synk-fixtures-go-mvs/d"

func GetB() string {
	return "b:1.0.0"
}

func GetD() string {
	return d.GetD()
}
