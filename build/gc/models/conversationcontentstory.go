package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontentstoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontentstoryDud struct { 
    


    


    

}

// Conversationcontentstory - Story object.
type Conversationcontentstory struct { 
    // VarType - Type of ephemeral story attachment.
    VarType string `json:"type"`


    // Url - URL to the ephemeral story.
    Url string `json:"url"`


    // ReplyToId - ID of the ephemeral story being replied to.
    ReplyToId string `json:"replyToId"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentstory) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontentstory) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontentstory

    if ConversationcontentstoryMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontentstoryMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Url string `json:"url"`
        
        ReplyToId string `json:"replyToId"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

