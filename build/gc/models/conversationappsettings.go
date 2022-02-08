package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationappsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationappsettingsDud struct { 
    


    

}

// Conversationappsettings - Conversation settings that handles chats within the messenger
type Conversationappsettings struct { 
    // ShowAgentTypingIndicator - The toggle to enable or disable typing indicator for messenger
    ShowAgentTypingIndicator bool `json:"showAgentTypingIndicator"`


    // ShowUserTypingIndicator - The toggle to enable or disable typing indicator for messenger
    ShowUserTypingIndicator bool `json:"showUserTypingIndicator"`

}

// String returns a JSON representation of the model
func (o *Conversationappsettings) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationappsettings) MarshalJSON() ([]byte, error) {
    type Alias Conversationappsettings

    if ConversationappsettingsMarshalled {
        return []byte("{}"), nil
    }
    ConversationappsettingsMarshalled = true

    return json.Marshal(&struct { 
        ShowAgentTypingIndicator bool `json:"showAgentTypingIndicator"`
        
        ShowUserTypingIndicator bool `json:"showUserTypingIndicator"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}
