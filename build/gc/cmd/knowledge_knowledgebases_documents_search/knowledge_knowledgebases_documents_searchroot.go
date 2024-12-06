package knowledge_knowledgebases_documents_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_search_suggestions"
)

func init() {
	knowledge_knowledgebases_documents_searchCmd.AddCommand(knowledge_knowledgebases_documents_search_suggestions.Cmdknowledge_knowledgebases_documents_search_suggestions())
	knowledge_knowledgebases_documents_searchCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_documents_searchCmd.Short, knowledge_knowledgebases_documents_search_suggestions.Description, )
	knowledge_knowledgebases_documents_searchCmd.Long = knowledge_knowledgebases_documents_searchCmd.Short
}