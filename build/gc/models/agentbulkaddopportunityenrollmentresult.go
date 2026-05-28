package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbulkaddopportunityenrollmentresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbulkaddopportunityenrollmentresultDud struct { 
    


    


    

}

// Agentbulkaddopportunityenrollmentresult
type Agentbulkaddopportunityenrollmentresult struct { 
    // Status - The status indicating the result of the bulk operation for this item
    Status string `json:"status"`


    // VarError - The error result if the operation failed
    VarError Bulkopportunitieserror `json:"error"`


    // Enrollment - The enrollment result that was added
    Enrollment Agentbulkaddopportunityenrollment `json:"enrollment"`

}

// String returns a JSON representation of the model
func (o *Agentbulkaddopportunityenrollmentresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbulkaddopportunityenrollmentresult) MarshalJSON() ([]byte, error) {
    type Alias Agentbulkaddopportunityenrollmentresult

    if AgentbulkaddopportunityenrollmentresultMarshalled {
        return []byte("{}"), nil
    }
    AgentbulkaddopportunityenrollmentresultMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarError Bulkopportunitieserror `json:"error"`
        
        Enrollment Agentbulkaddopportunityenrollment `json:"enrollment"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

