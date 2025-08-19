package knowledge_knowledgebases_export

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_export_jobs"
)

func init() {
	knowledge_knowledgebases_exportCmd.AddCommand(knowledge_knowledgebases_export_jobs.Cmdknowledge_knowledgebases_export_jobs())
	knowledge_knowledgebases_exportCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_exportCmd.Short, knowledge_knowledgebases_export_jobs.Description, )
	knowledge_knowledgebases_exportCmd.Long = knowledge_knowledgebases_exportCmd.Short
}
