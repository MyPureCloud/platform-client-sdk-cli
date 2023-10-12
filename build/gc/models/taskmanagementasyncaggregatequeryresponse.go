package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementasyncaggregatequeryresponseDud struct { 
    


    

}

// Taskmanagementasyncaggregatequeryresponse
type Taskmanagementasyncaggregatequeryresponse struct { 
    // Results
    Results []Taskmanagementaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementasyncaggregatequeryresponse) String() string {
     o.Results = []Taskmanagementaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementasyncaggregatequeryresponse

    if TaskmanagementasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Taskmanagementaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Taskmanagementaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

