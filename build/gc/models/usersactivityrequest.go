package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersactivityrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersactivityrequestDud struct { 
    

}

// Usersactivityrequest
type Usersactivityrequest struct { 
    // UserIds - The IDs of the users for whom to fetch their current activity state
    UserIds []string `json:"userIds"`

}

// String returns a JSON representation of the model
func (o *Usersactivityrequest) String() string {
     o.UserIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersactivityrequest) MarshalJSON() ([]byte, error) {
    type Alias Usersactivityrequest

    if UsersactivityrequestMarshalled {
        return []byte("{}"), nil
    }
    UsersactivityrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserIds []string `json:"userIds"`
        *Alias
    }{

        
        UserIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

