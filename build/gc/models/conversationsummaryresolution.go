package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummaryresolutionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummaryresolutionDud struct { 
    


    


    Confidence float32 `json:"confidence"`


    Outcome string `json:"outcome"`

}

// Conversationsummaryresolution
type Conversationsummaryresolution struct { 
    // Text - The text of the resolution.
    Text string `json:"text"`


    // Description - The description.
    Description string `json:"description"`


    


    

}

// String returns a JSON representation of the model
func (o *Conversationsummaryresolution) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummaryresolution) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummaryresolution

    if ConversationsummaryresolutionMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummaryresolutionMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

