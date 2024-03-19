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
    // Enabled - The toggle to enable or disable conversations
    Enabled bool `json:"enabled"`


    // ShowAgentTypingIndicator - The toggle to enable or disable typing indicator for messenger
    ShowAgentTypingIndicator bool `json:"showAgentTypingIndicator"`


    // ShowUserTypingIndicator - The toggle to enable or disable typing indicator for messenger
    ShowUserTypingIndicator bool `json:"showUserTypingIndicator"`


    // AutoStart - The auto start for the messenger conversation
    AutoStart Autostart `json:"autoStart"`


    // Markdown - The markdown for the messenger app
    Markdown Markdown `json:"markdown"`


    // ConversationDisconnect - The conversation disconnect settings for the messenger app
    ConversationDisconnect Conversationdisconnectsettings `json:"conversationDisconnect"`


    // ConversationClear - The conversation clear settings for the messenger app
    ConversationClear Conversationclearsettings `json:"conversationClear"`


    // Humanize - The humanize conversations settings for the messenger app
    Humanize Humanize `json:"humanize"`

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
        
        Enabled bool `json:"enabled"`
        
        ShowAgentTypingIndicator bool `json:"showAgentTypingIndicator"`
        
        ShowUserTypingIndicator bool `json:"showUserTypingIndicator"`
        
        AutoStart Autostart `json:"autoStart"`
        
        Markdown Markdown `json:"markdown"`
        
        ConversationDisconnect Conversationdisconnectsettings `json:"conversationDisconnect"`
        
        ConversationClear Conversationclearsettings `json:"conversationClear"`
        
        Humanize Humanize `json:"humanize"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

