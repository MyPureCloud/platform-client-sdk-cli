package externalcontacts_merge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_merge", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/merge")
	externalcontacts_mergeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_merge"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_mergeCmd)
}

func Cmdexternalcontacts_merge() *cobra.Command {
	return externalcontacts_mergeCmd
}
