package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VideomediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VideomediasettingsDud struct { 
    

}

// Videomediasettings
type Videomediasettings struct { 
    // Deployment - The web messaging deployment that video sessions for this queue originate from.
    Deployment Addressableentityref `json:"deployment"`

}

// String returns a JSON representation of the model
func (o *Videomediasettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Videomediasettings) MarshalJSON() ([]byte, error) {
    type Alias Videomediasettings

    if VideomediasettingsMarshalled {
        return []byte("{}"), nil
    }
    VideomediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        Deployment Addressableentityref `json:"deployment"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

