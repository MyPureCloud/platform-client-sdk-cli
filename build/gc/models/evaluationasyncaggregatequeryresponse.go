package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationasyncaggregatequeryresponseDud struct { 
    


    

}

// Evaluationasyncaggregatequeryresponse
type Evaluationasyncaggregatequeryresponse struct { 
    // Results
    Results []Evaluationaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Evaluationasyncaggregatequeryresponse) String() string {
     o.Results = []Evaluationaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluationasyncaggregatequeryresponse

    if EvaluationasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Evaluationaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Evaluationaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

