package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationpushfailedmessagereferencesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationpushfailedmessagereferencesDud struct { 
    

}

// Conversationpushfailedmessagereferences - References of failed messages requiring a push.
type Conversationpushfailedmessagereferences struct { 
    // Id - Unique ID of the message.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Conversationpushfailedmessagereferences) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationpushfailedmessagereferences) MarshalJSON() ([]byte, error) {
    type Alias Conversationpushfailedmessagereferences

    if ConversationpushfailedmessagereferencesMarshalled {
        return []byte("{}"), nil
    }
    ConversationpushfailedmessagereferencesMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

