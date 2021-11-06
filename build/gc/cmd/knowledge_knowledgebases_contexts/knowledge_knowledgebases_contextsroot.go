package knowledge_knowledgebases_contexts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_contexts_values"
)

func init() {
	knowledge_knowledgebases_contextsCmd.AddCommand(knowledge_knowledgebases_contexts_values.Cmdknowledge_knowledgebases_contexts_values())
	knowledge_knowledgebases_contextsCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_contextsCmd.Short, knowledge_knowledgebases_contexts_values.Description, )
	knowledge_knowledgebases_contextsCmd.Long = knowledge_knowledgebases_contextsCmd.Short
}
