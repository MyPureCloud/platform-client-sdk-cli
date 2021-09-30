package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotaggregatequeryresponseDud struct { 
    

}

// Botaggregatequeryresponse
type Botaggregatequeryresponse struct { 
    // Results
    Results []Botaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Botaggregatequeryresponse) String() string {
    
    
     o.Results = []Botaggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Botaggregatequeryresponse

    if BotaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    BotaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Botaggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Botaggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

