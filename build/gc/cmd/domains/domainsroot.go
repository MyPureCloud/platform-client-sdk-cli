
package domains

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/feedback"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/versions"
)

func init() {
	domainsCmd.AddCommand(feedback.Cmdfeedback())
	domainsCmd.AddCommand(versions.Cmdversions())
}
