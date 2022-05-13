package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkspacememberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkspacememberDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workspacemember
type Workspacemember struct { 
    


    // Name
    Name string `json:"name"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // MemberType - The workspace member type.
    MemberType string `json:"memberType"`


    // Member
    Member Domainentityref `json:"member"`


    // User
    User User `json:"user"`


    // Group
    Group Group `json:"group"`


    // SecurityProfile
    SecurityProfile Securityprofile `json:"securityProfile"`


    

}

// String returns a JSON representation of the model
func (o *Workspacemember) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workspacemember) MarshalJSON() ([]byte, error) {
    type Alias Workspacemember

    if WorkspacememberMarshalled {
        return []byte("{}"), nil
    }
    WorkspacememberMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Workspace Domainentityref `json:"workspace"`
        
        MemberType string `json:"memberType"`
        
        Member Domainentityref `json:"member"`
        
        User User `json:"user"`
        
        Group Group `json:"group"`
        
        SecurityProfile Securityprofile `json:"securityProfile"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

