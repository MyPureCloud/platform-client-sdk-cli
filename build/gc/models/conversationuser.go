package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationuserDud struct { 
    

}

// Conversationuser
type Conversationuser struct { 
    // Id - The globally unique identifier for this user.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Conversationuser) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationuser) MarshalJSON() ([]byte, error) {
    type Alias Conversationuser

    if ConversationuserMarshalled {
        return []byte("{}"), nil
    }
    ConversationuserMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

