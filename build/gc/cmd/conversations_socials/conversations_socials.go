package conversations_socials

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_socials", "SWAGGER_OVERRIDE_/api/v2/conversations/socials")
	conversations_socialsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_socials"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_socialsCmd)
}

func Cmdconversations_socials() *cobra.Command {
	return conversations_socialsCmd
}
