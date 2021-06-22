package authorization_divisionspermitted

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisionspermitted_paged"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisionspermitted_me"
)

func init() {
	authorization_divisionspermittedCmd.AddCommand(authorization_divisionspermitted_paged.Cmdauthorization_divisionspermitted_paged())
	authorization_divisionspermittedCmd.AddCommand(authorization_divisionspermitted_me.Cmdauthorization_divisionspermitted_me())
	authorization_divisionspermittedCmd.Short = utils.GenerateCustomDescription(authorization_divisionspermittedCmd.Short, authorization_divisionspermitted_paged.Description, authorization_divisionspermitted_me.Description, )
	authorization_divisionspermittedCmd.Long = authorization_divisionspermittedCmd.Short
}
