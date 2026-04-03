package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotasyncaggregatequeryresponseDud struct { 
    


    

}

// Copilotasyncaggregatequeryresponse
type Copilotasyncaggregatequeryresponse struct { 
    // Results
    Results []Copilotaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Copilotasyncaggregatequeryresponse) String() string {
     o.Results = []Copilotaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Copilotasyncaggregatequeryresponse

    if CopilotasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    CopilotasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Copilotaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Copilotaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

