package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagerDud struct { 
    


    Ref string `json:"$ref"`

}

// Manager - Defines a SCIM manager.
type Manager struct { 
    // Value - The ID of the manager.
    Value string `json:"value"`


    

}

// String returns a JSON representation of the model
func (o *Manager) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Manager) MarshalJSON() ([]byte, error) {
    type Alias Manager

    if ManagerMarshalled {
        return []byte("{}"), nil
    }
    ManagerMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

