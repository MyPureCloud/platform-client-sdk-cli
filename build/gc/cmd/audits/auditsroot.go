package audits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/audits_query"
)

func init() {
	auditsCmd.AddCommand(audits_query.Cmdaudits_query())
	auditsCmd.Short = utils.GenerateCustomDescription(auditsCmd.Short, audits_query.Description, )
	auditsCmd.Long = auditsCmd.Short
}
