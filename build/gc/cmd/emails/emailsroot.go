package emails

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/emails_settings"
)

func init() {
	emailsCmd.AddCommand(emails_settings.Cmdemails_settings())
	emailsCmd.Short = utils.GenerateCustomDescription(emailsCmd.Short, emails_settings.Description, )
	emailsCmd.Long = emailsCmd.Short
}
