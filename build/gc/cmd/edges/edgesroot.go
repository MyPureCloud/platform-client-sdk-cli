
package edges

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/didpools"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dids"
)

func init() {
	edgesCmd.AddCommand(didpools.Cmddidpools())
	edgesCmd.AddCommand(dids.Cmddids())
}
