package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasemanagementasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasemanagementasyncaggregatequeryresponseDud struct { 
    


    

}

// Casemanagementasyncaggregatequeryresponse
type Casemanagementasyncaggregatequeryresponse struct { 
    // Results
    Results []Casemanagementaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Casemanagementasyncaggregatequeryresponse) String() string {
     o.Results = []Casemanagementaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casemanagementasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Casemanagementasyncaggregatequeryresponse

    if CasemanagementasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    CasemanagementasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Casemanagementaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Casemanagementaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

