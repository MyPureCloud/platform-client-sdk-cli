package knowledge_knowledgebases_sources_servicenow

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_sources_servicenow_sync"
)

func init() {
	knowledge_knowledgebases_sources_servicenowCmd.AddCommand(knowledge_knowledgebases_sources_servicenow_sync.Cmdknowledge_knowledgebases_sources_servicenow_sync())
	knowledge_knowledgebases_sources_servicenowCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_sources_servicenowCmd.Short, knowledge_knowledgebases_sources_servicenow_sync.Description, )
	knowledge_knowledgebases_sources_servicenowCmd.Long = knowledge_knowledgebases_sources_servicenowCmd.Short
}
