package documentationfile

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("documentationfile", "SWAGGER_OVERRIDE_/api/v2/documentation")
	documentationfileCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("documentationfile"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(documentationfileCmd)
}

func Cmddocumentationfile() *cobra.Command {
	return documentationfileCmd
}
