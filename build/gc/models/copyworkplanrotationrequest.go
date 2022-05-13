package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopyworkplanrotationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopyworkplanrotationrequestDud struct { 
    

}

// Copyworkplanrotationrequest
type Copyworkplanrotationrequest struct { 
    // Name - Name to apply to the new copy of the work plan rotation
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Copyworkplanrotationrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copyworkplanrotationrequest) MarshalJSON() ([]byte, error) {
    type Alias Copyworkplanrotationrequest

    if CopyworkplanrotationrequestMarshalled {
        return []byte("{}"), nil
    }
    CopyworkplanrotationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

