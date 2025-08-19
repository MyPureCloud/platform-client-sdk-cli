package knowledge_knowledgebases_chunks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_chunks_search"
)

func init() {
	knowledge_knowledgebases_chunksCmd.AddCommand(knowledge_knowledgebases_chunks_search.Cmdknowledge_knowledgebases_chunks_search())
	knowledge_knowledgebases_chunksCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_chunksCmd.Short, knowledge_knowledgebases_chunks_search.Description, )
	knowledge_knowledgebases_chunksCmd.Long = knowledge_knowledgebases_chunksCmd.Short
}
