package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotaggregatequeryresponseDud struct { 
    

}

// Copilotaggregatequeryresponse
type Copilotaggregatequeryresponse struct { 
    // Results
    Results []Copilotaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Copilotaggregatequeryresponse) String() string {
     o.Results = []Copilotaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Copilotaggregatequeryresponse

    if CopilotaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    CopilotaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Copilotaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Copilotaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

