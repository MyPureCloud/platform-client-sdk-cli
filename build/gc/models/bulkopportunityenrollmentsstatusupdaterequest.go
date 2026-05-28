package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkopportunityenrollmentsstatusupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkopportunityenrollmentsstatusupdaterequestDud struct { 
    


    


    

}

// Bulkopportunityenrollmentsstatusupdaterequest
type Bulkopportunityenrollmentsstatusupdaterequest struct { 
    // EnrollmentIds - The IDs of the enrollments to update
    EnrollmentIds []string `json:"enrollmentIds"`


    // Status - The status to set for all enrollments specified in this request
    Status string `json:"status"`


    // ReviewNote - Supervisor's note explaining the agent's enrollment status change
    ReviewNote string `json:"reviewNote"`

}

// String returns a JSON representation of the model
func (o *Bulkopportunityenrollmentsstatusupdaterequest) String() string {
     o.EnrollmentIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkopportunityenrollmentsstatusupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkopportunityenrollmentsstatusupdaterequest

    if BulkopportunityenrollmentsstatusupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    BulkopportunityenrollmentsstatusupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        EnrollmentIds []string `json:"enrollmentIds"`
        
        Status string `json:"status"`
        
        ReviewNote string `json:"reviewNote"`
        *Alias
    }{

        
        EnrollmentIds: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

