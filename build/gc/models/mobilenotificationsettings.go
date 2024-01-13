package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MobilenotificationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MobilenotificationsettingsDud struct { 
    

}

// Mobilenotificationsettings - Settings for a user's mobile notifications
type Mobilenotificationsettings struct { 
    // When - When the user should receive notifications
    When string `json:"when"`

}

// String returns a JSON representation of the model
func (o *Mobilenotificationsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mobilenotificationsettings) MarshalJSON() ([]byte, error) {
    type Alias Mobilenotificationsettings

    if MobilenotificationsettingsMarshalled {
        return []byte("{}"), nil
    }
    MobilenotificationsettingsMarshalled = true

    return json.Marshal(&struct {
        
        When string `json:"when"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

