package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryopportunityenrollmentresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryopportunityenrollmentresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    

}

// Queryopportunityenrollmentresult
type Queryopportunityenrollmentresult struct { 
    


    // OpportunityId - The ID of the opportunity
    OpportunityId string `json:"opportunityId"`


    // Agent - A reference to the agent who created the enrollment
    Agent Userreference `json:"agent"`


    // Status - The current status of the enrollment
    Status string `json:"status"`


    // Schedule - The schedule on which the enrollment was added when this enrollment was approved
    Schedule Buschedulereference `json:"schedule"`


    // SystemMessageCode - The system-generated message code about enrollment processing results or failures
    SystemMessageCode string `json:"systemMessageCode"`


    // ReviewNote - Supervisor's note explaining the agent's enrollment status change
    ReviewNote string `json:"reviewNote"`


    // DenialCode - The denial code
    DenialCode string `json:"denialCode"`


    // Metadata - The metadata for the enrollment
    Metadata Queryopportunityenrollmentmetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Queryopportunityenrollmentresult) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryopportunityenrollmentresult) MarshalJSON() ([]byte, error) {
    type Alias Queryopportunityenrollmentresult

    if QueryopportunityenrollmentresultMarshalled {
        return []byte("{}"), nil
    }
    QueryopportunityenrollmentresultMarshalled = true

    return json.Marshal(&struct {
        
        OpportunityId string `json:"opportunityId"`
        
        Agent Userreference `json:"agent"`
        
        Status string `json:"status"`
        
        Schedule Buschedulereference `json:"schedule"`
        
        SystemMessageCode string `json:"systemMessageCode"`
        
        ReviewNote string `json:"reviewNote"`
        
        DenialCode string `json:"denialCode"`
        
        Metadata Queryopportunityenrollmentmetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

