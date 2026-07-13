package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationaccessattributesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationaccessattributesresponseDud struct { 
    

}

// Conversationaccessattributesresponse
type Conversationaccessattributesresponse struct { 
    // AccessAttributes - The attributes that define which users may access a conversation
    AccessAttributes []string `json:"accessAttributes"`

}

// String returns a JSON representation of the model
func (o *Conversationaccessattributesresponse) String() string {
     o.AccessAttributes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationaccessattributesresponse) MarshalJSON() ([]byte, error) {
    type Alias Conversationaccessattributesresponse

    if ConversationaccessattributesresponseMarshalled {
        return []byte("{}"), nil
    }
    ConversationaccessattributesresponseMarshalled = true

    return json.Marshal(&struct {
        
        AccessAttributes []string `json:"accessAttributes"`
        *Alias
    }{

        
        AccessAttributes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

