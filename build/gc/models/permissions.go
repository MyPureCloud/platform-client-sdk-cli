package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PermissionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PermissionsDud struct { 
    

}

// Permissions
type Permissions struct { 
    // Ids - List of permission ids.
    Ids []string `json:"ids"`

}

// String returns a JSON representation of the model
func (o *Permissions) String() string {
     o.Ids = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Permissions) MarshalJSON() ([]byte, error) {
    type Alias Permissions

    if PermissionsMarshalled {
        return []byte("{}"), nil
    }
    PermissionsMarshalled = true

    return json.Marshal(&struct {
        
        Ids []string `json:"ids"`
        *Alias
    }{

        
        Ids: []string{""},
        

        Alias: (*Alias)(u),
    })
}

