package telephony_providers_edges_mediastatistics_conversations_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_mediastatistics_conversations_communications_mediaresources"
)

func init() {
	telephony_providers_edges_mediastatistics_conversations_communicationsCmd.AddCommand(telephony_providers_edges_mediastatistics_conversations_communications_mediaresources.Cmdtelephony_providers_edges_mediastatistics_conversations_communications_mediaresources())
	telephony_providers_edges_mediastatistics_conversations_communicationsCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_mediastatistics_conversations_communicationsCmd.Short, telephony_providers_edges_mediastatistics_conversations_communications_mediaresources.Description, )
	telephony_providers_edges_mediastatistics_conversations_communicationsCmd.Long = telephony_providers_edges_mediastatistics_conversations_communicationsCmd.Short
}
