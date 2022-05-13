package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkspaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkspaceDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workspace
type Workspace struct { 
    


    // Name - The current name of the workspace.
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // IsCurrentUserWorkspace
    IsCurrentUserWorkspace bool `json:"isCurrentUserWorkspace"`


    // User
    User Domainentityref `json:"user"`


    // Bucket
    Bucket string `json:"bucket"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // Summary
    Summary Workspacesummary `json:"summary"`


    // Acl
    Acl []string `json:"acl"`


    // Description
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Workspace) String() string {
    
    
    
    
    
    
    
    
     o.Acl = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workspace) MarshalJSON() ([]byte, error) {
    type Alias Workspace

    if WorkspaceMarshalled {
        return []byte("{}"), nil
    }
    WorkspaceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        IsCurrentUserWorkspace bool `json:"isCurrentUserWorkspace"`
        
        User Domainentityref `json:"user"`
        
        Bucket string `json:"bucket"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Summary Workspacesummary `json:"summary"`
        
        Acl []string `json:"acl"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Acl: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

