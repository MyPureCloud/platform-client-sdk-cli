
package responsemgt

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/libraries"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responses"
)

func init() {
	responsemgtCmd.AddCommand(libraries.Cmdlibraries())
	responsemgtCmd.AddCommand(responses.Cmdresponses())
}
