package knowledge_knowledgebases_parse

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_parse_jobs"
)

func init() {
	knowledge_knowledgebases_parseCmd.AddCommand(knowledge_knowledgebases_parse_jobs.Cmdknowledge_knowledgebases_parse_jobs())
	knowledge_knowledgebases_parseCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_parseCmd.Short, knowledge_knowledgebases_parse_jobs.Description, )
	knowledge_knowledgebases_parseCmd.Long = knowledge_knowledgebases_parseCmd.Short
}
