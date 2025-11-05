package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationschemadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationschemadataDud struct { 
    


    


    

}

// Conversationschemadata
type Conversationschemadata struct { 
    // Schema - Schema that defines attributes.
    Schema Conversationschemareference `json:"schema"`


    // Attributes - Attributes to use for filtering; number of elements: min: 1, max: 5.
    Attributes []Conversationschemaattribute `json:"attributes"`


    // Operator - Operator to apply for multiple attributes, default: MatchAll
    Operator string `json:"operator"`

}

// String returns a JSON representation of the model
func (o *Conversationschemadata) String() string {
    
     o.Attributes = []Conversationschemaattribute{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationschemadata) MarshalJSON() ([]byte, error) {
    type Alias Conversationschemadata

    if ConversationschemadataMarshalled {
        return []byte("{}"), nil
    }
    ConversationschemadataMarshalled = true

    return json.Marshal(&struct {
        
        Schema Conversationschemareference `json:"schema"`
        
        Attributes []Conversationschemaattribute `json:"attributes"`
        
        Operator string `json:"operator"`
        *Alias
    }{

        


        
        Attributes: []Conversationschemaattribute{{}},
        


        

        Alias: (*Alias)(u),
    })
}

