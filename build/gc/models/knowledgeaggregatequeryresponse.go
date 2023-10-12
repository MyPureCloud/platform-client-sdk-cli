package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeaggregatequeryresponseDud struct { 
    

}

// Knowledgeaggregatequeryresponse
type Knowledgeaggregatequeryresponse struct { 
    // Results
    Results []Knowledgeaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Knowledgeaggregatequeryresponse) String() string {
     o.Results = []Knowledgeaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeaggregatequeryresponse

    if KnowledgeaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Knowledgeaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Knowledgeaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

