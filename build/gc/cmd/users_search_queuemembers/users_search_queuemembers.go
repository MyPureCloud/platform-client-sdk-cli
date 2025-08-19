package users_search_queuemembers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_search_queuemembers", "SWAGGER_OVERRIDE_/api/v2/users/search/queuemembers")
	users_search_queuemembersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_search_queuemembers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_search_queuemembersCmd)
}

func Cmdusers_search_queuemembers() *cobra.Command {
	return users_search_queuemembersCmd
}
