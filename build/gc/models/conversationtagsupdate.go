package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationtagsupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationtagsupdateDud struct { 
    

}

// Conversationtagsupdate
type Conversationtagsupdate struct { 
    // ExternalTag - The external tag associated with the conversation.
    ExternalTag string `json:"externalTag"`

}

// String returns a JSON representation of the model
func (o *Conversationtagsupdate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationtagsupdate) MarshalJSON() ([]byte, error) {
    type Alias Conversationtagsupdate

    if ConversationtagsupdateMarshalled {
        return []byte("{}"), nil
    }
    ConversationtagsupdateMarshalled = true

    return json.Marshal(&struct { 
        ExternalTag string `json:"externalTag"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

