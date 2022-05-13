package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShareMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShareDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Share
type Share struct { 
    


    // Name
    Name string `json:"name"`


    // SharedEntityType
    SharedEntityType string `json:"sharedEntityType"`


    // SharedEntity
    SharedEntity Domainentityref `json:"sharedEntity"`


    // MemberType
    MemberType string `json:"memberType"`


    // Member
    Member Domainentityref `json:"member"`


    // SharedBy
    SharedBy Domainentityref `json:"sharedBy"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // User
    User User `json:"user"`


    // Group
    Group Group `json:"group"`


    

}

// String returns a JSON representation of the model
func (o *Share) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Share) MarshalJSON() ([]byte, error) {
    type Alias Share

    if ShareMarshalled {
        return []byte("{}"), nil
    }
    ShareMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SharedEntityType string `json:"sharedEntityType"`
        
        SharedEntity Domainentityref `json:"sharedEntity"`
        
        MemberType string `json:"memberType"`
        
        Member Domainentityref `json:"member"`
        
        SharedBy Domainentityref `json:"sharedBy"`
        
        Workspace Domainentityref `json:"workspace"`
        
        User User `json:"user"`
        
        Group Group `json:"group"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

