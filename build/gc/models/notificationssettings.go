package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationssettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationssettingsDud struct { 
    

}

// Notificationssettings - Notification settings that handles messenger notifications
type Notificationssettings struct { 
    // Enabled - The toggle to enable or disable notifications
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Notificationssettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationssettings) MarshalJSON() ([]byte, error) {
    type Alias Notificationssettings

    if NotificationssettingsMarshalled {
        return []byte("{}"), nil
    }
    NotificationssettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

