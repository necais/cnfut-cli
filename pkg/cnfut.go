package pkg

import "github.com/oktayalizada/cnfut/entities"

func Copy(srcDest entities.SourceDestination) {
	switch srcDest.SourceType {
	case "s3":
		if srcDest.DestinationType == "azure" {
			fromS3ToAzure(srcDest)
		} else if srcDest.DestinationType == "local" {
			fromS3ToLocal(srcDest)
		} else if srcDest.DestinationType == "google" {
			fromS3ToGoogle(srcDest)
		} else {
			fromS3ToS3(srcDest)
		}
	case "local":
		if srcDest.DestinationType == "azure" {
			fromLocalToAzure(srcDest)
		} else if srcDest.DestinationType == "local" {
			fromLocalToLocal(srcDest)
		} else if srcDest.DestinationType == "google" {
			fromLocalToS3(srcDest)
		} else {
			fromLocalToGoogle(srcDest)
		}
	case "azure":
		if srcDest.DestinationType == "azure" {
			fromAzureToAzure(srcDest)
		} else if srcDest.DestinationType == "local" {
			fromAzureToLocal(srcDest)
		} else if srcDest.DestinationType == "google" {
			fromLocalToGoogle(srcDest)
		} else {
			fromAzureToS3(srcDest)
		}
	case "google":
		if srcDest.DestinationType == "azure" {
			fromGoogleToAzure(srcDest)
		} else if srcDest.DestinationType == "local" {
			fromGoogleToLocal(srcDest)
		} else if srcDest.DestinationType == "google" {
			fromGoogleToGoogle(srcDest)
		} else {
			fromGoogleToS3(srcDest)
		}
	}
}
