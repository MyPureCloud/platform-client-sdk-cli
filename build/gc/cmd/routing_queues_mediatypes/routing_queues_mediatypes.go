package routing_queues_mediatypes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_queues_mediatypes", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/mediatypes")
	routing_queues_mediatypesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_queues_mediatypes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_queues_mediatypesCmd)
}

func Cmdrouting_queues_mediatypes() *cobra.Command {
	return routing_queues_mediatypesCmd
}
