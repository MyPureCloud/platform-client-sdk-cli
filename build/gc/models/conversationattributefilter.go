package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationattributefilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationattributefilterDud struct { 
    

}

// Conversationattributefilter
type Conversationattributefilter struct { 
    // Schemas - Schemas and attributes to filter for.
    Schemas []Conversationschemadata `json:"schemas"`

}

// String returns a JSON representation of the model
func (o *Conversationattributefilter) String() string {
     o.Schemas = []Conversationschemadata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationattributefilter) MarshalJSON() ([]byte, error) {
    type Alias Conversationattributefilter

    if ConversationattributefilterMarshalled {
        return []byte("{}"), nil
    }
    ConversationattributefilterMarshalled = true

    return json.Marshal(&struct {
        
        Schemas []Conversationschemadata `json:"schemas"`
        *Alias
    }{

        
        Schemas: []Conversationschemadata{{}},
        

        Alias: (*Alias)(u),
    })
}

