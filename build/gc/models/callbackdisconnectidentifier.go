package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallbackdisconnectidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallbackdisconnectidentifierDud struct { 
    


    

}

// Callbackdisconnectidentifier
type Callbackdisconnectidentifier struct { 
    // ConversationId - The Conversation Id.
    ConversationId string `json:"conversationId"`


    // CallbackId - The callback id.
    CallbackId string `json:"callbackId"`

}

// String returns a JSON representation of the model
func (o *Callbackdisconnectidentifier) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callbackdisconnectidentifier) MarshalJSON() ([]byte, error) {
    type Alias Callbackdisconnectidentifier

    if CallbackdisconnectidentifierMarshalled {
        return []byte("{}"), nil
    }
    CallbackdisconnectidentifierMarshalled = true

    return json.Marshal(&struct { 
        ConversationId string `json:"conversationId"`
        
        CallbackId string `json:"callbackId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

