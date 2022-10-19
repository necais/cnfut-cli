package cnfut

import (
	"github.com/oktayalizada/cnfut/middleware"
)

type SourceDestination struct {
	Source          string
	Destination     string
	SourceType      string
	DestinationType string
}

func Copy(source, sourceType, destination, destinationType string, concurrent bool) {
	if middleware.CheckInput(source, sourceType, destination, destinationType) {
		test()
	}
}

func test() {

}