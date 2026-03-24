package knowledge_sources_synchronizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_sources_synchronizations_uploads"
)

func init() {
	knowledge_sources_synchronizationsCmd.AddCommand(knowledge_sources_synchronizations_uploads.Cmdknowledge_sources_synchronizations_uploads())
	knowledge_sources_synchronizationsCmd.Short = utils.GenerateCustomDescription(knowledge_sources_synchronizationsCmd.Short, knowledge_sources_synchronizations_uploads.Description, )
	knowledge_sources_synchronizationsCmd.Long = knowledge_sources_synchronizationsCmd.Short
}
