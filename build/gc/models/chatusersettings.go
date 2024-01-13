package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatusersettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatusersettingsDud struct { 
    

}

// Chatusersettings - Settings for a chat user
type Chatusersettings struct { 
    // Mobile - Settings for mobile devices
    Mobile Mobilesettings `json:"mobile"`

}

// String returns a JSON representation of the model
func (o *Chatusersettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatusersettings) MarshalJSON() ([]byte, error) {
    type Alias Chatusersettings

    if ChatusersettingsMarshalled {
        return []byte("{}"), nil
    }
    ChatusersettingsMarshalled = true

    return json.Marshal(&struct {
        
        Mobile Mobilesettings `json:"mobile"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

