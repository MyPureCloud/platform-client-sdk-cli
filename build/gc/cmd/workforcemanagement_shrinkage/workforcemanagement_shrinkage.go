package workforcemanagement_shrinkage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_shrinkage", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/shrinkage")
	workforcemanagement_shrinkageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_shrinkage"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_shrinkageCmd)
}

func Cmdworkforcemanagement_shrinkage() *cobra.Command {
	return workforcemanagement_shrinkageCmd
}
