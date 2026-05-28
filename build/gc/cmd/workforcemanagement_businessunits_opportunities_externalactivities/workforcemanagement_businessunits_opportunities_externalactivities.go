package workforcemanagement_businessunits_opportunities_externalactivities

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_opportunities_externalactivities", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/opportunities/externalactivities")
	workforcemanagement_businessunits_opportunities_externalactivitiesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_opportunities_externalactivities"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_opportunities_externalactivitiesCmd)
}

func Cmdworkforcemanagement_businessunits_opportunities_externalactivities() *cobra.Command {
	return workforcemanagement_businessunits_opportunities_externalactivitiesCmd
}
