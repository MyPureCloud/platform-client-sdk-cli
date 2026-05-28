package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpportunityresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpportunityresultDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Opportunityresult
type Opportunityresult struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // StartDate - The start date and time of the opportunity in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end date and time of the opportunity in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // Status - The current status of the opportunity
    Status string `json:"status"`


    // OpenDate - The date and time when the opportunity opens for enrollment in ISO-8601 format. If not provided or in the past, it will be automatically updated to the current time when the opportunity is published
    OpenDate time.Time `json:"openDate"`


    // DeadlineDate - The deadline date and time for enrollment in the opportunity in ISO-8601 format
    DeadlineDate time.Time `json:"deadlineDate"`


    // Name - The name of the opportunity
    Name string `json:"name"`


    // Description - Additional details describing the purpose or context of this opportunity
    Description string `json:"description"`


    // ActivityCodeId - The ID of the activity code associated with the opportunity
    ActivityCodeId string `json:"activityCodeId"`


    // ApprovalType - The approval type for enrollments
    ApprovalType string `json:"approvalType"`


    // AgentCount - The total number of agents invited to this opportunity
    AgentCount int `json:"agentCount"`


    // Capacity - The maximum capacity (enrollment slots) for this opportunity
    Capacity int `json:"capacity"`


    // EnrollmentProcessingCount - The number of enrollments currently being processed
    EnrollmentProcessingCount int `json:"enrollmentProcessingCount"`


    // EnrollmentCounts - The counts for enrollment statuses
    EnrollmentCounts Opportunityenrollmentcounts `json:"enrollmentCounts"`


    // PublishedDate - The date and time when the opportunity was published in ISO-8601 format
    PublishedDate time.Time `json:"publishedDate"`


    // ClosedDate - The date and time when the opportunity was closed in ISO-8601 format
    ClosedDate time.Time `json:"closedDate"`


    // SystemMessageCode - The system-generated message code about opportunity processing issues or validation failures
    SystemMessageCode string `json:"systemMessageCode"`


    // Metadata - The metadata for the opportunity
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Opportunityresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opportunityresult) MarshalJSON() ([]byte, error) {
    type Alias Opportunityresult

    if OpportunityresultMarshalled {
        return []byte("{}"), nil
    }
    OpportunityresultMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        Status string `json:"status"`
        
        OpenDate time.Time `json:"openDate"`
        
        DeadlineDate time.Time `json:"deadlineDate"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        ApprovalType string `json:"approvalType"`
        
        AgentCount int `json:"agentCount"`
        
        Capacity int `json:"capacity"`
        
        EnrollmentProcessingCount int `json:"enrollmentProcessingCount"`
        
        EnrollmentCounts Opportunityenrollmentcounts `json:"enrollmentCounts"`
        
        PublishedDate time.Time `json:"publishedDate"`
        
        ClosedDate time.Time `json:"closedDate"`
        
        SystemMessageCode string `json:"systemMessageCode"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

