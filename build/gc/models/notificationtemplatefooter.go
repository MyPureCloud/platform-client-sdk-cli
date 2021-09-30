package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationtemplatefooterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationtemplatefooterDud struct { 
    

}

// Notificationtemplatefooter - Template footer object.
type Notificationtemplatefooter struct { 
    // Text - Footer text. For WhatsApp, ignored.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplatefooter) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationtemplatefooter) MarshalJSON() ([]byte, error) {
    type Alias Notificationtemplatefooter

    if NotificationtemplatefooterMarshalled {
        return []byte("{}"), nil
    }
    NotificationtemplatefooterMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

