package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchopportunityrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchopportunityrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Patchopportunityrequest
type Patchopportunityrequest struct { 
    // StartDate - The start date and time of the opportunity in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date and time of the opportunity in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // OpenDate - The date and time when the opportunity opens for enrollment in ISO-8601 format. If not provided or in the past, it will be automatically updated to the current time when the opportunity is published
    OpenDate Valuewrapperinstant `json:"openDate"`


    // DeadlineDate - The deadline date and time for enrollment in the opportunity in ISO-8601 format
    DeadlineDate time.Time `json:"deadlineDate"`


    // Name - The name of the opportunity
    Name string `json:"name"`


    // Description - Additional details describing the purpose or context of this opportunity
    Description Valuewrapperstring `json:"description"`


    // ActivityCodeId - The ID of the activity code associated with the opportunity
    ActivityCodeId string `json:"activityCodeId"`


    // ApprovalType - The approval type for enrollments
    ApprovalType string `json:"approvalType"`


    // Capacity - The maximum capacity (enrollment slots) for this opportunity
    Capacity int `json:"capacity"`


    // AgentIds - The IDs of the agents that are invited to the opportunity
    AgentIds Listwrapperstring `json:"agentIds"`


    // Metadata - The metadata for the opportunity
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Patchopportunityrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchopportunityrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchopportunityrequest

    if PatchopportunityrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchopportunityrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        OpenDate Valuewrapperinstant `json:"openDate"`
        
        DeadlineDate time.Time `json:"deadlineDate"`
        
        Name string `json:"name"`
        
        Description Valuewrapperstring `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        ApprovalType string `json:"approvalType"`
        
        Capacity int `json:"capacity"`
        
        AgentIds Listwrapperstring `json:"agentIds"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

