package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleopaqueidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleopaqueidDud struct { 
    

}

// Appleopaqueid
type Appleopaqueid struct { 
    // Value
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Appleopaqueid) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleopaqueid) MarshalJSON() ([]byte, error) {
    type Alias Appleopaqueid

    if AppleopaqueidMarshalled {
        return []byte("{}"), nil
    }
    AppleopaqueidMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

