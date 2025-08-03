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
    // Enabled - The toggle to enable or disable notifications. When enabled, PushNotificationTitle and PushNotificationBody localization keys are required.
    Enabled bool `json:"enabled"`


    // NotificationContentType - The notification content type settings for messenger
    NotificationContentType string `json:"notificationContentType"`

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
        
        NotificationContentType string `json:"notificationContentType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

