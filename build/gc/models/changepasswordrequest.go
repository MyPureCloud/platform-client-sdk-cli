package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChangepasswordrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChangepasswordrequestDud struct { 
    

}

// Changepasswordrequest
type Changepasswordrequest struct { 
    // NewPassword - The new password
    NewPassword string `json:"newPassword"`

}

// String returns a JSON representation of the model
func (o *Changepasswordrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Changepasswordrequest) MarshalJSON() ([]byte, error) {
    type Alias Changepasswordrequest

    if ChangepasswordrequestMarshalled {
        return []byte("{}"), nil
    }
    ChangepasswordrequestMarshalled = true

    return json.Marshal(&struct {
        
        NewPassword string `json:"newPassword"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

