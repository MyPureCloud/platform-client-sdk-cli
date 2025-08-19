package contentmanagement_workspaces_tagvalues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_workspaces_tagvalues_query"
)

func init() {
	contentmanagement_workspaces_tagvaluesCmd.AddCommand(contentmanagement_workspaces_tagvalues_query.Cmdcontentmanagement_workspaces_tagvalues_query())
	contentmanagement_workspaces_tagvaluesCmd.Short = utils.GenerateCustomDescription(contentmanagement_workspaces_tagvaluesCmd.Short, contentmanagement_workspaces_tagvalues_query.Description, )
	contentmanagement_workspaces_tagvaluesCmd.Long = contentmanagement_workspaces_tagvaluesCmd.Short
}
