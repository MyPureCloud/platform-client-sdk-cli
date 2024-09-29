package knowledge_knowledgebases_sources_salesforce

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_sources_salesforce_sync"
)

func init() {
	knowledge_knowledgebases_sources_salesforceCmd.AddCommand(knowledge_knowledgebases_sources_salesforce_sync.Cmdknowledge_knowledgebases_sources_salesforce_sync())
	knowledge_knowledgebases_sources_salesforceCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_sources_salesforceCmd.Short, knowledge_knowledgebases_sources_salesforce_sync.Description, )
	knowledge_knowledgebases_sources_salesforceCmd.Long = knowledge_knowledgebases_sources_salesforceCmd.Short
}
