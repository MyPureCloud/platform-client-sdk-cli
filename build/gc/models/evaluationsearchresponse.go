package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsearchresponseDud struct { 
    PageSize int `json:"pageSize"`


    PageNumber int `json:"pageNumber"`


    Results []Evaluationresponse `json:"results"`


    Aggregations map[string]Evaluationsearchaggregationresponse `json:"aggregations"`

}

// Evaluationsearchresponse
type Evaluationsearchresponse struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Evaluationsearchresponse) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsearchresponse

    if EvaluationsearchresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsearchresponseMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

