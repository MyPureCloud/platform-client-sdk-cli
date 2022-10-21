package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeammembersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeammembersDud struct { 
    

}

// Teammembers
type Teammembers struct { 
    // MemberIds - A list of the ids of the members to add or remove
    MemberIds []string `json:"memberIds"`

}

// String returns a JSON representation of the model
func (o *Teammembers) String() string {
     o.MemberIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teammembers) MarshalJSON() ([]byte, error) {
    type Alias Teammembers

    if TeammembersMarshalled {
        return []byte("{}"), nil
    }
    TeammembersMarshalled = true

    return json.Marshal(&struct {
        
        MemberIds []string `json:"memberIds"`
        *Alias
    }{

        
        MemberIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

