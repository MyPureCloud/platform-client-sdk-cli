package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrappershifttradetargetrequestitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrappershifttradetargetrequestitemDud struct { 
    

}

// Valuewrappershifttradetargetrequestitem
type Valuewrappershifttradetargetrequestitem struct { 
    // Value - The value for the associated field
    Value Shifttradetargetrequestitem `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrappershifttradetargetrequestitem) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrappershifttradetargetrequestitem) MarshalJSON() ([]byte, error) {
    type Alias Valuewrappershifttradetargetrequestitem

    if ValuewrappershifttradetargetrequestitemMarshalled {
        return []byte("{}"), nil
    }
    ValuewrappershifttradetargetrequestitemMarshalled = true

    return json.Marshal(&struct {
        
        Value Shifttradetargetrequestitem `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

