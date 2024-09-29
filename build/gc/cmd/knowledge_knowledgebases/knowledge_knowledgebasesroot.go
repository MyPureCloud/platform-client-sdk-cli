package knowledge_knowledgebases

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_categories"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_languages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_unanswered"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_labels"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_export"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_import"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_sources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_uploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_synchronize"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_operations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_parse"
)

func init() {
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_categories.Cmdknowledge_knowledgebases_categories())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_search.Cmdknowledge_knowledgebases_search())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_languages.Cmdknowledge_knowledgebases_languages())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_unanswered.Cmdknowledge_knowledgebases_unanswered())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_labels.Cmdknowledge_knowledgebases_labels())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_documents.Cmdknowledge_knowledgebases_documents())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_export.Cmdknowledge_knowledgebases_export())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_import.Cmdknowledge_knowledgebases_import())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_sources.Cmdknowledge_knowledgebases_sources())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_uploads.Cmdknowledge_knowledgebases_uploads())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_synchronize.Cmdknowledge_knowledgebases_synchronize())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_operations.Cmdknowledge_knowledgebases_operations())
	knowledge_knowledgebasesCmd.AddCommand(knowledge_knowledgebases_parse.Cmdknowledge_knowledgebases_parse())
	knowledge_knowledgebasesCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebasesCmd.Short, knowledge_knowledgebases_categories.Description, knowledge_knowledgebases_search.Description, knowledge_knowledgebases_languages.Description, knowledge_knowledgebases_unanswered.Description, knowledge_knowledgebases_labels.Description, knowledge_knowledgebases_documents.Description, knowledge_knowledgebases_export.Description, knowledge_knowledgebases_import.Description, knowledge_knowledgebases_sources.Description, knowledge_knowledgebases_uploads.Description, knowledge_knowledgebases_synchronize.Description, knowledge_knowledgebases_operations.Description, knowledge_knowledgebases_parse.Description, )
	knowledge_knowledgebasesCmd.Long = knowledge_knowledgebasesCmd.Short
}
