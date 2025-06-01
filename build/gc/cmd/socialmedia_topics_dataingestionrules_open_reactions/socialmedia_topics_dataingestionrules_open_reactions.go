package socialmedia_topics_dataingestionrules_open_reactions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_open_reactions", "SWAGGER_OVERRIDE_/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{ruleId}/reactions")
	socialmedia_topics_dataingestionrules_open_reactionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_topics_dataingestionrules_open_reactions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_topics_dataingestionrules_open_reactionsCmd)
}

func Cmdsocialmedia_topics_dataingestionrules_open_reactions() *cobra.Command {
	return socialmedia_topics_dataingestionrules_open_reactionsCmd
}
