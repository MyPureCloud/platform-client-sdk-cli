package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentbulkstatusupdateopportunityenrollmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentbulkstatusupdateopportunityenrollmentsrequestDud struct { 
    


    

}

// Agentbulkstatusupdateopportunityenrollmentsrequest
type Agentbulkstatusupdateopportunityenrollmentsrequest struct { 
    // EnrollmentIds - The IDs of the enrollments to update
    EnrollmentIds []string `json:"enrollmentIds"`


    // Status - The status to set for all enrollments specified in this request
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Agentbulkstatusupdateopportunityenrollmentsrequest) String() string {
     o.EnrollmentIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentbulkstatusupdateopportunityenrollmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Agentbulkstatusupdateopportunityenrollmentsrequest

    if AgentbulkstatusupdateopportunityenrollmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    AgentbulkstatusupdateopportunityenrollmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        EnrollmentIds []string `json:"enrollmentIds"`
        
        Status string `json:"status"`
        *Alias
    }{

        
        EnrollmentIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

