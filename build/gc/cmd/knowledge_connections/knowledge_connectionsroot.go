package knowledge_connections

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_connections_options"
)

func init() {
	knowledge_connectionsCmd.AddCommand(knowledge_connections_options.Cmdknowledge_connections_options())
	knowledge_connectionsCmd.Short = utils.GenerateCustomDescription(knowledge_connectionsCmd.Short, knowledge_connections_options.Description, )
	knowledge_connectionsCmd.Long = knowledge_connectionsCmd.Short
}
