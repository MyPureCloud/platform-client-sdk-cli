package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasemanagementaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasemanagementaggregatequeryresponseDud struct { 
    

}

// Casemanagementaggregatequeryresponse
type Casemanagementaggregatequeryresponse struct { 
    // Results
    Results []Casemanagementaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Casemanagementaggregatequeryresponse) String() string {
     o.Results = []Casemanagementaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casemanagementaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Casemanagementaggregatequeryresponse

    if CasemanagementaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    CasemanagementaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Casemanagementaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Casemanagementaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

