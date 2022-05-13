package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimuserroleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimuserroleDud struct { 
    

}

// Scimuserrole - Defines a user role.
type Scimuserrole struct { 
    // Value - The role of the Genesys Cloud user.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Scimuserrole) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimuserrole) MarshalJSON() ([]byte, error) {
    type Alias Scimuserrole

    if ScimuserroleMarshalled {
        return []byte("{}"), nil
    }
    ScimuserroleMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

