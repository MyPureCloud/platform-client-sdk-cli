package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotificationtemplateparameterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotificationtemplateparameterDud struct { 
    


    

}

// Notificationtemplateparameter - Template parameters for placeholders in template.
type Notificationtemplateparameter struct { 
    // Name - Parameter name.
    Name string `json:"name"`


    // Text - Parameter text value.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplateparameter) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notificationtemplateparameter) MarshalJSON() ([]byte, error) {
    type Alias Notificationtemplateparameter

    if NotificationtemplateparameterMarshalled {
        return []byte("{}"), nil
    }
    NotificationtemplateparameterMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Text string `json:"text"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

