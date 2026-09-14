package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationattributeDud struct { 
    


    

}

// Conversationattribute
type Conversationattribute struct { 
    // Schema - The Conversation Custom Attributes schema that the variable is bound to.
    Schema Conversationattributeschema `json:"schema"`


    // AttributeName - The name of the attribute within the schema that the variable value is bound to.
    AttributeName string `json:"attributeName"`

}

// String returns a JSON representation of the model
func (o *Conversationattribute) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationattribute) MarshalJSON() ([]byte, error) {
    type Alias Conversationattribute

    if ConversationattributeMarshalled {
        return []byte("{}"), nil
    }
    ConversationattributeMarshalled = true

    return json.Marshal(&struct {
        
        Schema Conversationattributeschema `json:"schema"`
        
        AttributeName string `json:"attributeName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

