package pkg

import "fmt"

type SourceDestination struct {
	Source          string
	Destination     string
	SourceType      string
	DestinationType string
	Concurrent      bool
}

func Copy(srcDest SourceDestination) {
	fmt.Println("worked")
}

func test() {

}
