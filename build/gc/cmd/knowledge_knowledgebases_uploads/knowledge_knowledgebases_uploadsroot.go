package knowledge_knowledgebases_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_uploads_urls"
)

func init() {
	knowledge_knowledgebases_uploadsCmd.AddCommand(knowledge_knowledgebases_uploads_urls.Cmdknowledge_knowledgebases_uploads_urls())
	knowledge_knowledgebases_uploadsCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_uploadsCmd.Short, knowledge_knowledgebases_uploads_urls.Description, )
	knowledge_knowledgebases_uploadsCmd.Long = knowledge_knowledgebases_uploadsCmd.Short
}
