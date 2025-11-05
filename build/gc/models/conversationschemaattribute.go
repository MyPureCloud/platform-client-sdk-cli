package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationschemaattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationschemaattributeDud struct { 
    

}

// Conversationschemaattribute
type Conversationschemaattribute struct { 
    // Name - Name of the attribute as defined in the schema.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Conversationschemaattribute) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationschemaattribute) MarshalJSON() ([]byte, error) {
    type Alias Conversationschemaattribute

    if ConversationschemaattributeMarshalled {
        return []byte("{}"), nil
    }
    ConversationschemaattributeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

