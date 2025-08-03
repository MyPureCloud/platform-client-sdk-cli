package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagesubtypesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagesubtypesettingsDud struct { 
    

}

// Messagesubtypesettings
type Messagesubtypesettings struct { 
    // EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
    EnableAutoAnswer bool `json:"enableAutoAnswer"`

}

// String returns a JSON representation of the model
func (o *Messagesubtypesettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagesubtypesettings) MarshalJSON() ([]byte, error) {
    type Alias Messagesubtypesettings

    if MessagesubtypesettingsMarshalled {
        return []byte("{}"), nil
    }
    MessagesubtypesettingsMarshalled = true

    return json.Marshal(&struct {
        
        EnableAutoAnswer bool `json:"enableAutoAnswer"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

