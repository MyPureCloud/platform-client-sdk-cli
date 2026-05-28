package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbulkaddopportunityenrollmentsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbulkaddopportunityenrollmentsresponseDud struct { 
    


    

}

// Agentbulkaddopportunityenrollmentsresponse
type Agentbulkaddopportunityenrollmentsresponse struct { 
    // Results - The result for each requested item
    Results []Agentbulkaddopportunityenrollmentresult `json:"results"`


    // ErrorCount - The count of failed operations in the bulk request
    ErrorCount int `json:"errorCount"`

}

// String returns a JSON representation of the model
func (o *Agentbulkaddopportunityenrollmentsresponse) String() string {
     o.Results = []Agentbulkaddopportunityenrollmentresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbulkaddopportunityenrollmentsresponse) MarshalJSON() ([]byte, error) {
    type Alias Agentbulkaddopportunityenrollmentsresponse

    if AgentbulkaddopportunityenrollmentsresponseMarshalled {
        return []byte("{}"), nil
    }
    AgentbulkaddopportunityenrollmentsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Agentbulkaddopportunityenrollmentresult `json:"results"`
        
        ErrorCount int `json:"errorCount"`
        *Alias
    }{

        
        Results: []Agentbulkaddopportunityenrollmentresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

