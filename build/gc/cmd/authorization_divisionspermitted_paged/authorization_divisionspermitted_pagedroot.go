package authorization_divisionspermitted_paged

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_divisionspermitted_paged_me"
)

func init() {
	authorization_divisionspermitted_pagedCmd.AddCommand(authorization_divisionspermitted_paged_me.Cmdauthorization_divisionspermitted_paged_me())
	authorization_divisionspermitted_pagedCmd.Short = utils.GenerateCustomDescription(authorization_divisionspermitted_pagedCmd.Short, authorization_divisionspermitted_paged_me.Description, )
	authorization_divisionspermitted_pagedCmd.Long = authorization_divisionspermitted_pagedCmd.Short
}
