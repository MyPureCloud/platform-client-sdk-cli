package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PublishdraftinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PublishdraftinputDud struct { 
    

}

// Publishdraftinput - Draft to be published
type Publishdraftinput struct { 
    // Version - The current draft version.
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Publishdraftinput) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Publishdraftinput) MarshalJSON() ([]byte, error) {
    type Alias Publishdraftinput

    if PublishdraftinputMarshalled {
        return []byte("{}"), nil
    }
    PublishdraftinputMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

