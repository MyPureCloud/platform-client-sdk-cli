package workforcemanagement_businessunits_shifttrading_unmatched_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_unmatched_search", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/shifttrading/unmatched/search")
	workforcemanagement_businessunits_shifttrading_unmatched_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_unmatched_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_shifttrading_unmatched_searchCmd)
}

func Cmdworkforcemanagement_businessunits_shifttrading_unmatched_search() *cobra.Command {
	return workforcemanagement_businessunits_shifttrading_unmatched_searchCmd
}
