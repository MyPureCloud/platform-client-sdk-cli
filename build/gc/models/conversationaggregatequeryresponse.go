package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationaggregatequeryresponseDud struct { 
    

}

// Conversationaggregatequeryresponse
type Conversationaggregatequeryresponse struct { 
    // Results
    Results []Conversationaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Conversationaggregatequeryresponse) String() string {
    
    
     o.Results = []Conversationaggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Conversationaggregatequeryresponse

    if ConversationaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    ConversationaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Conversationaggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Conversationaggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

