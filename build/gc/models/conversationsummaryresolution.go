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
    


    


    


    

}

// Conversationsummaryresolution
type Conversationsummaryresolution struct { 
    // Confidence - The AI confidence value.
    Confidence float32 `json:"confidence"`


    // Text - The text of the resolution.
    Text string `json:"text"`


    // Description - The description.
    Description string `json:"description"`


    // Outcome - The outcome of the conversation's resolution.
    Outcome string `json:"outcome"`

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
        
        Confidence float32 `json:"confidence"`
        
        Text string `json:"text"`
        
        Description string `json:"description"`
        
        Outcome string `json:"outcome"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

