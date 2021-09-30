package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationthreadingwindowsettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationthreadingwindowsettingDud struct { 
    


    

}

// Conversationthreadingwindowsetting
type Conversationthreadingwindowsetting struct { 
    // MessengerType - The type of messenger
    MessengerType string `json:"messengerType"`


    // TimeoutInMinutes - The conversation threading window timeout (Minutes) of specified messenger type
    TimeoutInMinutes int `json:"timeoutInMinutes"`

}

// String returns a JSON representation of the model
func (o *Conversationthreadingwindowsetting) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationthreadingwindowsetting) MarshalJSON() ([]byte, error) {
    type Alias Conversationthreadingwindowsetting

    if ConversationthreadingwindowsettingMarshalled {
        return []byte("{}"), nil
    }
    ConversationthreadingwindowsettingMarshalled = true

    return json.Marshal(&struct { 
        MessengerType string `json:"messengerType"`
        
        TimeoutInMinutes int `json:"timeoutInMinutes"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

