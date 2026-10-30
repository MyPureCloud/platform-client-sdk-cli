package knowledge_sources_files

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_sources_files_counts"
)

func init() {
	knowledge_sources_filesCmd.AddCommand(knowledge_sources_files_counts.Cmdknowledge_sources_files_counts())
	knowledge_sources_filesCmd.Short = utils.GenerateCustomDescription(knowledge_sources_filesCmd.Short, knowledge_sources_files_counts.Description, )
	knowledge_sources_filesCmd.Long = knowledge_sources_filesCmd.Short
}
