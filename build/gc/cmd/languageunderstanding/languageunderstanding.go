package languageunderstanding

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("languageunderstanding", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding")
	languageunderstandingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstandingCmd)
}

func Cmdlanguageunderstanding() *cobra.Command {
	return languageunderstandingCmd
}
