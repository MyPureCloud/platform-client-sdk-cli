package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationaggregatequeryresponseDud struct { 
    

}

// Evaluationaggregatequeryresponse
type Evaluationaggregatequeryresponse struct { 
    // Results
    Results []Evaluationaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryresponse) String() string {
     o.Results = []Evaluationaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluationaggregatequeryresponse

    if EvaluationaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Evaluationaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Evaluationaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

