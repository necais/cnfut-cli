package middleware

import "github.com/oktayalizada/cnfut/helper"

var SupportedSystems = []string{"azure", "local", "s3"}

func CheckInput(source, sourceType, destination, destinationType string) bool {
	return helper.Contains(SupportedSystems, sourceType) && helper.Contains(SupportedSystems, destinationType)
}
