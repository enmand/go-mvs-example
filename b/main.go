package b

import "github.com/enmand/go-mvs-example/d"

func GetB() string {
	return "b:1.0.0"
}

func GetD() string {
	return d.GetD()
}
