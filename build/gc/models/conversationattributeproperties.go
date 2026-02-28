package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationattributepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationattributepropertiesDud struct { 
    


    

}

// Conversationattributeproperties
type Conversationattributeproperties struct { 
    // Schema - Schema that defines attributes.
    Schema Conversationschemareference `json:"schema"`


    // Name - Attribute name.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Conversationattributeproperties) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationattributeproperties) MarshalJSON() ([]byte, error) {
    type Alias Conversationattributeproperties

    if ConversationattributepropertiesMarshalled {
        return []byte("{}"), nil
    }
    ConversationattributepropertiesMarshalled = true

    return json.Marshal(&struct {
        
        Schema Conversationschemareference `json:"schema"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

