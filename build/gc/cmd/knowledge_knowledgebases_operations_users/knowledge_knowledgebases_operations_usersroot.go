package knowledge_knowledgebases_operations_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_operations_users_query"
)

func init() {
	knowledge_knowledgebases_operations_usersCmd.AddCommand(knowledge_knowledgebases_operations_users_query.Cmdknowledge_knowledgebases_operations_users_query())
	knowledge_knowledgebases_operations_usersCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_operations_usersCmd.Short, knowledge_knowledgebases_operations_users_query.Description, )
	knowledge_knowledgebases_operations_usersCmd.Long = knowledge_knowledgebases_operations_usersCmd.Short
}
