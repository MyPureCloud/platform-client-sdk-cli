package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcontenttextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcontenttextDud struct { 
    


    

}

// Conversationcontenttext - Message content element containing text only.
type Conversationcontenttext struct { 
    // VarType - Type of text content (Deprecated).
    VarType string `json:"type"`


    // Body - Text to be shown for this content element.
    Body string `json:"body"`

}

// String returns a JSON representation of the model
func (o *Conversationcontenttext) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcontenttext) MarshalJSON() ([]byte, error) {
    type Alias Conversationcontenttext

    if ConversationcontenttextMarshalled {
        return []byte("{}"), nil
    }
    ConversationcontenttextMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Body string `json:"body"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

