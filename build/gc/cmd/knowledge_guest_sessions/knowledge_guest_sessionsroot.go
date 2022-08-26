package knowledge_guest_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest_sessions_documents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest_sessions_categories"
)

func init() {
	knowledge_guest_sessionsCmd.AddCommand(knowledge_guest_sessions_documents.Cmdknowledge_guest_sessions_documents())
	knowledge_guest_sessionsCmd.AddCommand(knowledge_guest_sessions_categories.Cmdknowledge_guest_sessions_categories())
	knowledge_guest_sessionsCmd.Short = utils.GenerateCustomDescription(knowledge_guest_sessionsCmd.Short, knowledge_guest_sessions_documents.Description, knowledge_guest_sessions_categories.Description, )
	knowledge_guest_sessionsCmd.Long = knowledge_guest_sessionsCmd.Short
}
