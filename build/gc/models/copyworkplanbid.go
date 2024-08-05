package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopyworkplanbidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopyworkplanbidDud struct { 
    

}

// Copyworkplanbid
type Copyworkplanbid struct { 
    // Name - The name of the new work plan bid
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Copyworkplanbid) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copyworkplanbid) MarshalJSON() ([]byte, error) {
    type Alias Copyworkplanbid

    if CopyworkplanbidMarshalled {
        return []byte("{}"), nil
    }
    CopyworkplanbidMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

