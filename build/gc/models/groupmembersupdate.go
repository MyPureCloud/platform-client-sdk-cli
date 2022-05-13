package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupmembersupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupmembersupdateDud struct { 
    


    

}

// Groupmembersupdate
type Groupmembersupdate struct { 
    // MemberIds - A list of the ids of the members to add.
    MemberIds []string `json:"memberIds"`


    // Version - The current group version.
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Groupmembersupdate) String() string {
     o.MemberIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupmembersupdate) MarshalJSON() ([]byte, error) {
    type Alias Groupmembersupdate

    if GroupmembersupdateMarshalled {
        return []byte("{}"), nil
    }
    GroupmembersupdateMarshalled = true

    return json.Marshal(&struct {
        
        MemberIds []string `json:"memberIds"`
        
        Version int `json:"version"`
        *Alias
    }{

        
        MemberIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

