package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignusersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignusersDud struct { 
    


    

}

// Assignusers
type Assignusers struct { 
    // MembersToAssign - List of user ids to assign to a performance profile
    MembersToAssign []string `json:"membersToAssign"`


    // MembersToRemove - List of user ids to remove from a performance profile
    MembersToRemove []string `json:"membersToRemove"`

}

// String returns a JSON representation of the model
func (o *Assignusers) String() string {
    
    
     o.MembersToAssign = []string{""} 
    
    
    
     o.MembersToRemove = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignusers) MarshalJSON() ([]byte, error) {
    type Alias Assignusers

    if AssignusersMarshalled {
        return []byte("{}"), nil
    }
    AssignusersMarshalled = true

    return json.Marshal(&struct { 
        MembersToAssign []string `json:"membersToAssign"`
        
        MembersToRemove []string `json:"membersToRemove"`
        
        *Alias
    }{
        

        
        MembersToAssign: []string{""},
        

        

        
        MembersToRemove: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

