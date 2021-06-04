package audits_query_realtime

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/audits_query_realtime_servicemapping"
)

func init() {
	audits_query_realtimeCmd.AddCommand(audits_query_realtime_servicemapping.Cmdaudits_query_realtime_servicemapping())
	audits_query_realtimeCmd.Short = utils.GenerateCustomDescription(audits_query_realtimeCmd.Short, audits_query_realtime_servicemapping.Description, )
	audits_query_realtimeCmd.Long = audits_query_realtimeCmd.Short
}
