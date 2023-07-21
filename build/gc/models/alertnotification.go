package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertnotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertnotificationDud struct { 
    


    


    

}

// Alertnotification
type Alertnotification struct { 
    // Recipient - The entity to receive the notification.
    Recipient string `json:"recipient"`


    // NotificationTypes - The notification types the user will receive.
    NotificationTypes []string `json:"notificationTypes"`


    // Locale - The locale whose language will be used when sending alerts.  Locale should be in theformat language_COUNTRY where language is always lower case and country is always upper case.
    Locale string `json:"locale"`

}

// String returns a JSON representation of the model
func (o *Alertnotification) String() string {
    
     o.NotificationTypes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertnotification) MarshalJSON() ([]byte, error) {
    type Alias Alertnotification

    if AlertnotificationMarshalled {
        return []byte("{}"), nil
    }
    AlertnotificationMarshalled = true

    return json.Marshal(&struct {
        
        Recipient string `json:"recipient"`
        
        NotificationTypes []string `json:"notificationTypes"`
        
        Locale string `json:"locale"`
        *Alias
    }{

        


        
        NotificationTypes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

