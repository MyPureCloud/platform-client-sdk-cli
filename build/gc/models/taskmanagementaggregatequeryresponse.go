package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementaggregatequeryresponseDud struct { 
    

}

// Taskmanagementaggregatequeryresponse
type Taskmanagementaggregatequeryresponse struct { 
    // Results
    Results []Taskmanagementaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementaggregatequeryresponse) String() string {
     o.Results = []Taskmanagementaggregatedatacontainer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementaggregatequeryresponse

    if TaskmanagementaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Taskmanagementaggregatedatacontainer `json:"results"`
        *Alias
    }{

        
        Results: []Taskmanagementaggregatedatacontainer{{}},
        

        Alias: (*Alias)(u),
    })
}

