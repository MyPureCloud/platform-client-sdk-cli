package knowledge_guest_sessions_documents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest_sessions_documents_search"
)

func init() {
	knowledge_guest_sessions_documentsCmd.AddCommand(knowledge_guest_sessions_documents_search.Cmdknowledge_guest_sessions_documents_search())
	knowledge_guest_sessions_documentsCmd.Short = utils.GenerateCustomDescription(knowledge_guest_sessions_documentsCmd.Short, knowledge_guest_sessions_documents_search.Description, )
	knowledge_guest_sessions_documentsCmd.Long = knowledge_guest_sessions_documentsCmd.Short
}
