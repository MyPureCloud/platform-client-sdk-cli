
package subjects

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/bulkadd"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/bulkremove"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/bulkreplace"
)

func init() {
	subjectsCmd.AddCommand(bulkadd.Cmdbulkadd())
	subjectsCmd.AddCommand(bulkremove.Cmdbulkremove())
	subjectsCmd.AddCommand(bulkreplace.Cmdbulkreplace())
}
