package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MobilesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MobilesettingsDud struct { 
    

}

// Mobilesettings - Settings for mobile devices
type Mobilesettings struct { 
    // Notifications - Settings for a user's mobile notifications
    Notifications Mobilenotificationsettings `json:"notifications"`

}

// String returns a JSON representation of the model
func (o *Mobilesettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mobilesettings) MarshalJSON() ([]byte, error) {
    type Alias Mobilesettings

    if MobilesettingsMarshalled {
        return []byte("{}"), nil
    }
    MobilesettingsMarshalled = true

    return json.Marshal(&struct {
        
        Notifications Mobilenotificationsettings `json:"notifications"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

