package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationencryptionconfigurationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationencryptionconfigurationlistingDud struct { 
    


    


    

}

// Conversationencryptionconfigurationlisting
type Conversationencryptionconfigurationlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Conversationencryptionconfiguration `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Conversationencryptionconfigurationlisting) String() string {
    
     o.Entities = []Conversationencryptionconfiguration{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationencryptionconfigurationlisting) MarshalJSON() ([]byte, error) {
    type Alias Conversationencryptionconfigurationlisting

    if ConversationencryptionconfigurationlistingMarshalled {
        return []byte("{}"), nil
    }
    ConversationencryptionconfigurationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Conversationencryptionconfiguration `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Conversationencryptionconfiguration{{}},
        


        

        Alias: (*Alias)(u),
    })
}

