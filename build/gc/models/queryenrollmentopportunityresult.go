package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryenrollmentopportunityresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryenrollmentopportunityresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    

}

// Queryenrollmentopportunityresult
type Queryenrollmentopportunityresult struct { 
    


    // Name - The name of the opportunity
    Name string `json:"name"`


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


    // Capacity - The maximum capacity (enrollment slots) for this opportunity
    Capacity int `json:"capacity"`


    // EnrollmentCounts - The counts for enrollment statuses
    EnrollmentCounts Pendingandapprovedopportunityenrollmentcounts `json:"enrollmentCounts"`

}

// String returns a JSON representation of the model
func (o *Queryenrollmentopportunityresult) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryenrollmentopportunityresult) MarshalJSON() ([]byte, error) {
    type Alias Queryenrollmentopportunityresult

    if QueryenrollmentopportunityresultMarshalled {
        return []byte("{}"), nil
    }
    QueryenrollmentopportunityresultMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        DeadlineDate time.Time `json:"deadlineDate"`
        
        Status string `json:"status"`
        
        Capacity int `json:"capacity"`
        
        EnrollmentCounts Pendingandapprovedopportunityenrollmentcounts `json:"enrollmentCounts"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

