package greetings_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("greetings_groups", "SWAGGER_OVERRIDE_/api/v2/greetings/{greetingId}/groups")
	greetings_groupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("greetings_groups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(greetings_groupsCmd)
}

func Cmdgreetings_groups() *cobra.Command {
	return greetings_groupsCmd
}
