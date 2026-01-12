package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationuserdispositionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationuserdispositionDud struct { 
    


    


    

}

// Conversationuserdisposition
type Conversationuserdisposition struct { 
    // Code - User-defined wrap-up code for the conversation.
    Code string `json:"code"`


    // Notes - Text entered by the user to describe the call or disposition.
    Notes string `json:"notes"`


    // User - The user that wrapped up the conversation.
    User Addressableentityref `json:"user"`

}

// String returns a JSON representation of the model
func (o *Conversationuserdisposition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationuserdisposition) MarshalJSON() ([]byte, error) {
    type Alias Conversationuserdisposition

    if ConversationuserdispositionMarshalled {
        return []byte("{}"), nil
    }
    ConversationuserdispositionMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Notes string `json:"notes"`
        
        User Addressableentityref `json:"user"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

