package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsummaryreasonMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsummaryreasonDud struct { 
    


    


    

}

// Conversationsummaryreason
type Conversationsummaryreason struct { 
    // Confidence - The AI confidence value.
    Confidence float32 `json:"confidence"`


    // Text - The text of the reason.
    Text string `json:"text"`


    // Description - The description.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Conversationsummaryreason) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsummaryreason) MarshalJSON() ([]byte, error) {
    type Alias Conversationsummaryreason

    if ConversationsummaryreasonMarshalled {
        return []byte("{}"), nil
    }
    ConversationsummaryreasonMarshalled = true

    return json.Marshal(&struct {
        
        Confidence float32 `json:"confidence"`
        
        Text string `json:"text"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

