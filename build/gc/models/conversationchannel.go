package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationchannelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationchannelDud struct { 
    


    


    

}

// Conversationchannel
type Conversationchannel struct { 
    // VarType - The type or category of this channel.
    VarType string `json:"type"`


    // MessageType - Message type for messaging conversations.
    MessageType string `json:"messageType"`


    // Platform - The source provider for the conversation (e.g. Edge, PureCloud Messaging, PureCloud Email).
    Platform string `json:"platform"`

}

// String returns a JSON representation of the model
func (o *Conversationchannel) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationchannel) MarshalJSON() ([]byte, error) {
    type Alias Conversationchannel

    if ConversationchannelMarshalled {
        return []byte("{}"), nil
    }
    ConversationchannelMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        MessageType string `json:"messageType"`
        
        Platform string `json:"platform"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

