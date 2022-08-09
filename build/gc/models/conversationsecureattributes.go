package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsecureattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsecureattributesDud struct { 
    


    

}

// Conversationsecureattributes
type Conversationsecureattributes struct { 
    // Attributes - The map of attribute keys to values.
    Attributes map[string]string `json:"attributes"`


    // Version - The version used to detect conflicting updates when using PUT. Not used for PATCH.
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Conversationsecureattributes) String() string {
     o.Attributes = map[string]string{"": ""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsecureattributes) MarshalJSON() ([]byte, error) {
    type Alias Conversationsecureattributes

    if ConversationsecureattributesMarshalled {
        return []byte("{}"), nil
    }
    ConversationsecureattributesMarshalled = true

    return json.Marshal(&struct {
        
        Attributes map[string]string `json:"attributes"`
        
        Version int `json:"version"`
        *Alias
    }{

        
        Attributes: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

