package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationaccessattributesupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationaccessattributesupdateDud struct { 
    

}

// Conversationaccessattributesupdate
type Conversationaccessattributesupdate struct { 
    // AccessAttributes - The attributes that define which users may access a conversation. Values must be non-null, non-blank, and contain only alphanumeric characters, hyphens, and underscores.
    AccessAttributes []string `json:"accessAttributes"`

}

// String returns a JSON representation of the model
func (o *Conversationaccessattributesupdate) String() string {
     o.AccessAttributes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationaccessattributesupdate) MarshalJSON() ([]byte, error) {
    type Alias Conversationaccessattributesupdate

    if ConversationaccessattributesupdateMarshalled {
        return []byte("{}"), nil
    }
    ConversationaccessattributesupdateMarshalled = true

    return json.Marshal(&struct {
        
        AccessAttributes []string `json:"accessAttributes"`
        *Alias
    }{

        
        AccessAttributes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

