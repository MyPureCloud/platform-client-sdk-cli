package audits_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/audits_query_results"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/audits_query_realtime"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/audits_query_servicemapping"
)

func init() {
	audits_queryCmd.AddCommand(audits_query_results.Cmdaudits_query_results())
	audits_queryCmd.AddCommand(audits_query_realtime.Cmdaudits_query_realtime())
	audits_queryCmd.AddCommand(audits_query_servicemapping.Cmdaudits_query_servicemapping())
	audits_queryCmd.Short = utils.GenerateCustomDescription(audits_queryCmd.Short, audits_query_results.Description, audits_query_realtime.Description, audits_query_servicemapping.Description, )
	audits_queryCmd.Long = audits_queryCmd.Short
}
