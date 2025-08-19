package knowledge_knowledgebases_unanswered

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_unanswered_groups"
)

func init() {
	knowledge_knowledgebases_unansweredCmd.AddCommand(knowledge_knowledgebases_unanswered_groups.Cmdknowledge_knowledgebases_unanswered_groups())
	knowledge_knowledgebases_unansweredCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_unansweredCmd.Short, knowledge_knowledgebases_unanswered_groups.Description, )
	knowledge_knowledgebases_unansweredCmd.Long = knowledge_knowledgebases_unansweredCmd.Short
}
