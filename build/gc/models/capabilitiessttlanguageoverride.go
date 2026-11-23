package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiessttlanguageoverrideMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiessttlanguageoverrideDud struct { 
    

}

// Capabilitiessttlanguageoverride
type Capabilitiessttlanguageoverride struct { 
    // Engine
    Engine string `json:"engine"`

}

// String returns a JSON representation of the model
func (o *Capabilitiessttlanguageoverride) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiessttlanguageoverride) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiessttlanguageoverride

    if CapabilitiessttlanguageoverrideMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiessttlanguageoverrideMarshalled = true

    return json.Marshal(&struct {
        
        Engine string `json:"engine"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

