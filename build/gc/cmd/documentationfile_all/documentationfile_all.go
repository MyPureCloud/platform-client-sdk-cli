package documentationfile_all

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("documentationfile_all", "SWAGGER_OVERRIDE_/api/v2/documentation/all")
	documentationfile_allCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("documentationfile_all"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(documentationfile_allCmd)
}

func Cmddocumentationfile_all() *cobra.Command {
	return documentationfile_allCmd
}
