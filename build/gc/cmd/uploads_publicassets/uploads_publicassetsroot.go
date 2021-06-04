package uploads_publicassets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads_publicassets_images"
)

func init() {
	uploads_publicassetsCmd.AddCommand(uploads_publicassets_images.Cmduploads_publicassets_images())
	uploads_publicassetsCmd.Short = utils.GenerateCustomDescription(uploads_publicassetsCmd.Short, uploads_publicassets_images.Description, )
	uploads_publicassetsCmd.Long = uploads_publicassetsCmd.Short
}
