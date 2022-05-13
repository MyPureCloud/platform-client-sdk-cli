package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserobservationqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserobservationqueryresponseDud struct { 
    

}

// Userobservationqueryresponse
type Userobservationqueryresponse struct { 
    // Results
    Results []Userobservationdatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Userobservationqueryresponse) String() string {
     o.Results = []Userobservationdatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userobservationqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Userobservationqueryresponse

    if UserobservationqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    UserobservationqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Userobservationdatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Userobservationdatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

