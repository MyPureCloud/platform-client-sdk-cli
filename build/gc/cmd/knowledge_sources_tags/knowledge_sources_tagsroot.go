package knowledge_sources_tags

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_sources_tags_grouped"
)

func init() {
	knowledge_sources_tagsCmd.AddCommand(knowledge_sources_tags_grouped.Cmdknowledge_sources_tags_grouped())
	knowledge_sources_tagsCmd.Short = utils.GenerateCustomDescription(knowledge_sources_tagsCmd.Short, knowledge_sources_tags_grouped.Description, )
	knowledge_sources_tagsCmd.Long = knowledge_sources_tagsCmd.Short
}
