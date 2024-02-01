package knowledge_knowledgebases_operations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_operations_users"
)

func init() {
	knowledge_knowledgebases_operationsCmd.AddCommand(knowledge_knowledgebases_operations_users.Cmdknowledge_knowledgebases_operations_users())
	knowledge_knowledgebases_operationsCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_operationsCmd.Short, knowledge_knowledgebases_operations_users.Description, )
	knowledge_knowledgebases_operationsCmd.Long = knowledge_knowledgebases_operationsCmd.Short
}
