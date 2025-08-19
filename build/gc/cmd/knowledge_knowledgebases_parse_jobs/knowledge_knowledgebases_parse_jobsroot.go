package knowledge_knowledgebases_parse_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_parse_jobs_import"
)

func init() {
	knowledge_knowledgebases_parse_jobsCmd.AddCommand(knowledge_knowledgebases_parse_jobs_import.Cmdknowledge_knowledgebases_parse_jobs_import())
	knowledge_knowledgebases_parse_jobsCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_parse_jobsCmd.Short, knowledge_knowledgebases_parse_jobs_import.Description, )
	knowledge_knowledgebases_parse_jobsCmd.Long = knowledge_knowledgebases_parse_jobsCmd.Short
}
