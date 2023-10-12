package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeasyncaggregatequeryresponseDud struct { 
    


    

}

// Knowledgeasyncaggregatequeryresponse
type Knowledgeasyncaggregatequeryresponse struct { 
    // Results
    Results []Knowledgeaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Knowledgeasyncaggregatequeryresponse) String() string {
     o.Results = []Knowledgeaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeasyncaggregatequeryresponse

    if KnowledgeasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Knowledgeaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Knowledgeaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

