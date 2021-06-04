package learning_rules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("learning_rules", "SWAGGER_OVERRIDE_/api/v2/learning/rules")
	learning_rulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_rules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_rulesCmd)
}

func Cmdlearning_rules() *cobra.Command {
	return learning_rulesCmd
}
