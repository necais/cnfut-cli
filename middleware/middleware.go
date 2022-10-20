package middleware

import (
	"github.com/oktayalizada/cnfut/pkg"
	"github.com/oktayalizada/cnfut/utils"
)

var SupportedSystems = []string{"azure", "local", "s3"}

func CheckInput(srcDest pkg.SourceDestination) bool {
	return utils.Contains(SupportedSystems, srcDest.SourceType) && utils.Contains(SupportedSystems, srcDest.DestinationType)
}
