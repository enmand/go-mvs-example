package d

import "github.com/enmand/go-mvs-example/e"

func GetD() string {
	return "d:1.2.0"
}

func GetE() string {
	return e.GetE()
}
