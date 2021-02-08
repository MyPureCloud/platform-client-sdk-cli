
package queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/queueuser"
)

func init() {
	queuesCmd.AddCommand(queueuser.Cmdqueueuser())
}
