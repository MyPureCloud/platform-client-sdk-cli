package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueuemediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueuemediasettingsDud struct { 
    


    


    


    


    

}

// Queuemediasettings
type Queuemediasettings struct { 
    // Call - The queue media settings for call interactions.
    Call Mediasettings `json:"call"`


    // Callback - The queue media settings for callback interactions.
    Callback Callbackmediasettings `json:"callback"`


    // Chat - The queue media settings for chat interactions.
    Chat Mediasettings `json:"chat"`


    // Email - The queue media settings for email interactions.
    Email Mediasettings `json:"email"`


    // Message - The queue media settings for message interactions.
    Message Mediasettings `json:"message"`

}

// String returns a JSON representation of the model
func (o *Queuemediasettings) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queuemediasettings) MarshalJSON() ([]byte, error) {
    type Alias Queuemediasettings

    if QueuemediasettingsMarshalled {
        return []byte("{}"), nil
    }
    QueuemediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        Call Mediasettings `json:"call"`
        
        Callback Callbackmediasettings `json:"callback"`
        
        Chat Mediasettings `json:"chat"`
        
        Email Mediasettings `json:"email"`
        
        Message Mediasettings `json:"message"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

