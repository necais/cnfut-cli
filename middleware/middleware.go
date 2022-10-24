package middleware

import (
	"github.com/oktayalizada/cnfut/entities"
	"github.com/oktayalizada/cnfut/utils"
)

func ValidateAndFormat(srcDest *entities.SourceDestination) bool {
	return utils.Contains(entities.SupportedSystems, srcDest.SourceType) && utils.Contains(entities.SupportedSystems, srcDest.DestinationType)
}
