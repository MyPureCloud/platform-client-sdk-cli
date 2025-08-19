package socialmedia_topics_dataingestionrules_open_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_open_messages", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{ruleId}/messages")
	socialmedia_topics_dataingestionrules_open_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_open_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_topics_dataingestionrules_open_messagesCmd)
}

func Cmdsocialmedia_topics_dataingestionrules_open_messages() *cobra.Command {
	return socialmedia_topics_dataingestionrules_open_messagesCmd
}
