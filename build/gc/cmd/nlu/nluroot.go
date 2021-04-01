
package nlu

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/domains"
)

func init() {
	nluCmd.AddCommand(domains.Cmddomains())
}
