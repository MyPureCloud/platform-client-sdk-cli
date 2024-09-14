package knowledge_knowledgebases_documents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_answers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_variations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_copies"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_presentations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_feedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_views"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_bulk"
)

func init() {
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_answers.Cmdknowledge_knowledgebases_documents_answers())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_versions.Cmdknowledge_knowledgebases_documents_versions())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_variations.Cmdknowledge_knowledgebases_documents_variations())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_copies.Cmdknowledge_knowledgebases_documents_copies())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_presentations.Cmdknowledge_knowledgebases_documents_presentations())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_feedback.Cmdknowledge_knowledgebases_documents_feedback())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_query.Cmdknowledge_knowledgebases_documents_query())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_search.Cmdknowledge_knowledgebases_documents_search())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_views.Cmdknowledge_knowledgebases_documents_views())
	knowledge_knowledgebases_documentsCmd.AddCommand(knowledge_knowledgebases_documents_bulk.Cmdknowledge_knowledgebases_documents_bulk())
	knowledge_knowledgebases_documentsCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_documentsCmd.Short, knowledge_knowledgebases_documents_answers.Description, knowledge_knowledgebases_documents_versions.Description, knowledge_knowledgebases_documents_variations.Description, knowledge_knowledgebases_documents_copies.Description, knowledge_knowledgebases_documents_presentations.Description, knowledge_knowledgebases_documents_feedback.Description, knowledge_knowledgebases_documents_query.Description, knowledge_knowledgebases_documents_search.Description, knowledge_knowledgebases_documents_views.Description, knowledge_knowledgebases_documents_bulk.Description, )
	knowledge_knowledgebases_documentsCmd.Long = knowledge_knowledgebases_documentsCmd.Short
}