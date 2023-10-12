package analytics_knowledge

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_knowledge_aggregates"
)

func init() {
	analytics_knowledgeCmd.AddCommand(analytics_knowledge_aggregates.Cmdanalytics_knowledge_aggregates())
	analytics_knowledgeCmd.Short = utils.GenerateCustomDescription(analytics_knowledgeCmd.Short, analytics_knowledge_aggregates.Description, )
	analytics_knowledgeCmd.Long = analytics_knowledgeCmd.Short
}
