package webdeployments_token

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_token_refresh"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_token_oauthcodegrantjwtexchange"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_token_revoke"
)

func init() {
	webdeployments_tokenCmd.AddCommand(webdeployments_token_refresh.Cmdwebdeployments_token_refresh())
	webdeployments_tokenCmd.AddCommand(webdeployments_token_oauthcodegrantjwtexchange.Cmdwebdeployments_token_oauthcodegrantjwtexchange())
	webdeployments_tokenCmd.AddCommand(webdeployments_token_revoke.Cmdwebdeployments_token_revoke())
	webdeployments_tokenCmd.Short = utils.GenerateCustomDescription(webdeployments_tokenCmd.Short, webdeployments_token_refresh.Description, webdeployments_token_oauthcodegrantjwtexchange.Description, webdeployments_token_revoke.Description, )
	webdeployments_tokenCmd.Long = webdeployments_tokenCmd.Short
}
