package knowledge_knowledgebases_synchronize

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_synchronize_jobs"
)

func init() {
	knowledge_knowledgebases_synchronizeCmd.AddCommand(knowledge_knowledgebases_synchronize_jobs.Cmdknowledge_knowledgebases_synchronize_jobs())
	knowledge_knowledgebases_synchronizeCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_synchronizeCmd.Short, knowledge_knowledgebases_synchronize_jobs.Description, )
	knowledge_knowledgebases_synchronizeCmd.Long = knowledge_knowledgebases_synchronizeCmd.Short
}
