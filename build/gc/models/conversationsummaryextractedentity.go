package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummaryextractedentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummaryextractedentityDud struct { 
    


    

}

// Conversationsummaryextractedentity
type Conversationsummaryextractedentity struct { 
    // Label - The label for the extracted entity
    Label string `json:"label"`


    // Value - The value for the extracted entity
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Conversationsummaryextractedentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummaryextractedentity) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummaryextractedentity

    if ConversationsummaryextractedentityMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummaryextractedentityMarshalled = true

    return json.Marshal(&struct {
        
        Label string `json:"label"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

