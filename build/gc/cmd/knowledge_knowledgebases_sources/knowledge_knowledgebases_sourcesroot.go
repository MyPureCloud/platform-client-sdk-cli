package knowledge_knowledgebases_sources

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_sources_salesforce"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_sources_servicenow"
)

func init() {
	knowledge_knowledgebases_sourcesCmd.AddCommand(knowledge_knowledgebases_sources_salesforce.Cmdknowledge_knowledgebases_sources_salesforce())
	knowledge_knowledgebases_sourcesCmd.AddCommand(knowledge_knowledgebases_sources_servicenow.Cmdknowledge_knowledgebases_sources_servicenow())
	knowledge_knowledgebases_sourcesCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_sourcesCmd.Short, knowledge_knowledgebases_sources_salesforce.Description, knowledge_knowledgebases_sources_servicenow.Description, )
	knowledge_knowledgebases_sourcesCmd.Long = knowledge_knowledgebases_sourcesCmd.Short
}
