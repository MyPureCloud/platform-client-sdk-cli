package knowledge_knowledgebases_import

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_import_jobs"
)

func init() {
	knowledge_knowledgebases_importCmd.AddCommand(knowledge_knowledgebases_import_jobs.Cmdknowledge_knowledgebases_import_jobs())
	knowledge_knowledgebases_importCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_importCmd.Short, knowledge_knowledgebases_import_jobs.Description, )
	knowledge_knowledgebases_importCmd.Long = knowledge_knowledgebases_importCmd.Short
}
