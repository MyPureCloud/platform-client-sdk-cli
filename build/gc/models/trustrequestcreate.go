package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustrequestcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustrequestcreateDud struct { 
    


    

}

// Trustrequestcreate
type Trustrequestcreate struct { 
    // UserIds - The list of trustee users that are requesting access. If no users are specified, at least one group is required.
    UserIds []string `json:"userIds"`


    // GroupIds - The list of trustee groups that are requesting access. If no groups are specified, at least one user is required.
    GroupIds []string `json:"groupIds"`

}

// String returns a JSON representation of the model
func (o *Trustrequestcreate) String() string {
     o.UserIds = []string{""} 
     o.GroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustrequestcreate) MarshalJSON() ([]byte, error) {
    type Alias Trustrequestcreate

    if TrustrequestcreateMarshalled {
        return []byte("{}"), nil
    }
    TrustrequestcreateMarshalled = true

    return json.Marshal(&struct {
        
        UserIds []string `json:"userIds"`
        
        GroupIds []string `json:"groupIds"`
        *Alias
    }{

        
        UserIds: []string{""},
        


        
        GroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

