package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookpermissionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookpermissionDud struct { 
    Name string `json:"name"`


    VarType string `json:"type"`

}

// Facebookpermission - Facebook Permissions Model
type Facebookpermission struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Facebookpermission) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookpermission) MarshalJSON() ([]byte, error) {
    type Alias Facebookpermission

    if FacebookpermissionMarshalled {
        return []byte("{}"), nil
    }
    FacebookpermissionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

