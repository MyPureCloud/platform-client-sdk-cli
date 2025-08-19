package users_search_conversation

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_search_conversation", "SWAGGER_OVERRIDE_/api/v2/users/search/conversation")
	users_search_conversationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_search_conversation"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_search_conversationCmd)
}

func Cmdusers_search_conversation() *cobra.Command {
	return users_search_conversationCmd
}
