package knowledge_knowledgebases_uploads_urls

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_uploads_urls_jobs"
)

func init() {
	knowledge_knowledgebases_uploads_urlsCmd.AddCommand(knowledge_knowledgebases_uploads_urls_jobs.Cmdknowledge_knowledgebases_uploads_urls_jobs())
	knowledge_knowledgebases_uploads_urlsCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_uploads_urlsCmd.Short, knowledge_knowledgebases_uploads_urls_jobs.Description, )
	knowledge_knowledgebases_uploads_urlsCmd.Long = knowledge_knowledgebases_uploads_urlsCmd.Short
}
