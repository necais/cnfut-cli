package entities

type SourceDestination struct {
	Source          string
	Destination     string
	SourceType      string
	DestinationType string
	Concurrent      bool
	Region          string
}

var SupportedSystems = []string{"azure", "local", "s3"}

const DefaultAwsRegion = "eu-north-1"
