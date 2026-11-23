package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentasyncaggregatequeryresponseDud struct { 
    


    

}

// Agenticvirtualagentasyncaggregatequeryresponse
type Agenticvirtualagentasyncaggregatequeryresponse struct { 
    // Results
    Results []Agenticvirtualagentaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentasyncaggregatequeryresponse) String() string {
     o.Results = []Agenticvirtualagentaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentasyncaggregatequeryresponse

    if AgenticvirtualagentasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agenticvirtualagentaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Agenticvirtualagentaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

