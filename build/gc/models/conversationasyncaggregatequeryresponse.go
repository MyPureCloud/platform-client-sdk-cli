package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationasyncaggregatequeryresponseDud struct { 
    


    

}

// Conversationasyncaggregatequeryresponse
type Conversationasyncaggregatequeryresponse struct { 
    // Results
    Results []Conversationaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Conversationasyncaggregatequeryresponse) String() string {
     o.Results = []Conversationaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Conversationasyncaggregatequeryresponse

    if ConversationasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    ConversationasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Conversationaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Conversationaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

