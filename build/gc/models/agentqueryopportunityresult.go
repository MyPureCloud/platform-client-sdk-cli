package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentqueryopportunityresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentqueryopportunityresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    

}

// Agentqueryopportunityresult
type Agentqueryopportunityresult struct { 
    


    // Name - The name of the opportunity
    Name string `json:"name"`


    // Description - Additional details describing the purpose or context of this opportunity
    Description string `json:"description"`


    // ActivityCodeId - The ID of the activity code associated with the opportunity
    ActivityCodeId string `json:"activityCodeId"`


    // StartDate - The start date and time of the opportunity in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date and time of the opportunity in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // DeadlineDate - The deadline date and time for enrollment in the opportunity in ISO-8601 format
    DeadlineDate time.Time `json:"deadlineDate"`


    // Status - The current status of the opportunity
    Status string `json:"status"`


    // Capacity - The maximum capacity for this opportunity
    Capacity int `json:"capacity"`


    // EnrollmentCounts - Subset of enrollment counts which are relevant to the agent
    EnrollmentCounts Pendingandapprovedopportunityenrollmentcounts `json:"enrollmentCounts"`


    // Enrollment - The agent's enrollment in this opportunity, if enrolled
    Enrollment Agentopportunityenrollmentresult `json:"enrollment"`


    // Metadata - The metadata for the opportunity
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Agentqueryopportunityresult) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentqueryopportunityresult) MarshalJSON() ([]byte, error) {
    type Alias Agentqueryopportunityresult

    if AgentqueryopportunityresultMarshalled {
        return []byte("{}"), nil
    }
    AgentqueryopportunityresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        DeadlineDate time.Time `json:"deadlineDate"`
        
        Status string `json:"status"`
        
        Capacity int `json:"capacity"`
        
        EnrollmentCounts Pendingandapprovedopportunityenrollmentcounts `json:"enrollmentCounts"`
        
        Enrollment Agentopportunityenrollmentresult `json:"enrollment"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

