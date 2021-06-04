package contentmanagement_workspaces

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_workspaces_tagvalues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_workspaces_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement_workspaces_documents"
)

func init() {
	contentmanagement_workspacesCmd.AddCommand(contentmanagement_workspaces_tagvalues.Cmdcontentmanagement_workspaces_tagvalues())
	contentmanagement_workspacesCmd.AddCommand(contentmanagement_workspaces_members.Cmdcontentmanagement_workspaces_members())
	contentmanagement_workspacesCmd.AddCommand(contentmanagement_workspaces_documents.Cmdcontentmanagement_workspaces_documents())
	contentmanagement_workspacesCmd.Short = utils.GenerateCustomDescription(contentmanagement_workspacesCmd.Short, contentmanagement_workspaces_tagvalues.Description, contentmanagement_workspaces_members.Description, contentmanagement_workspaces_documents.Description, )
	contentmanagement_workspacesCmd.Long = contentmanagement_workspacesCmd.Short
}
