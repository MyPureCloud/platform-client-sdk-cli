package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateassignusersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateassignusersDud struct { 
    

}

// Validateassignusers
type Validateassignusers struct { 
    // MembersToAssign - List of user ids to assign to a performance profile
    MembersToAssign []string `json:"membersToAssign"`

}

// String returns a JSON representation of the model
func (o *Validateassignusers) String() string {
     o.MembersToAssign = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateassignusers) MarshalJSON() ([]byte, error) {
    type Alias Validateassignusers

    if ValidateassignusersMarshalled {
        return []byte("{}"), nil
    }
    ValidateassignusersMarshalled = true

    return json.Marshal(&struct {
        
        MembersToAssign []string `json:"membersToAssign"`
        *Alias
    }{

        
        MembersToAssign: []string{""},
        

        Alias: (*Alias)(u),
    })
}

