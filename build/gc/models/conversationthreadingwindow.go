package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationthreadingwindowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationthreadingwindowDud struct { 
    Id string `json:"id"`


    


    DefaultTimeoutMinutes int `json:"defaultTimeoutMinutes"`

}

// Conversationthreadingwindow
type Conversationthreadingwindow struct { 
    


    // Settings - The conversation threading window timeout (Minutes) for each messaging type
    Settings []Conversationthreadingwindowsetting `json:"settings"`


    

}

// String returns a JSON representation of the model
func (o *Conversationthreadingwindow) String() string {
    
    
    
    
     o.Settings = []Conversationthreadingwindowsetting{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationthreadingwindow) MarshalJSON() ([]byte, error) {
    type Alias Conversationthreadingwindow

    if ConversationthreadingwindowMarshalled {
        return []byte("{}"), nil
    }
    ConversationthreadingwindowMarshalled = true

    return json.Marshal(&struct { 
        
        
        Settings []Conversationthreadingwindowsetting `json:"settings"`
        
        
        
        *Alias
    }{
        

        

        

        
        Settings: []Conversationthreadingwindowsetting{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

