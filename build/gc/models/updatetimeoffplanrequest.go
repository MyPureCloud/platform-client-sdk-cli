package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatetimeoffplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatetimeoffplanrequestDud struct { 
    


    


    


    


    


    


    

}

// Updatetimeoffplanrequest
type Updatetimeoffplanrequest struct { 
    // Name - The name of this time off plan.
    Name string `json:"name"`


    // ActivityCodeIds - The set of activity code IDs to associate with this time off plan.
    ActivityCodeIds Setwrapperstring `json:"activityCodeIds"`


    // TimeOffLimitIds - The set of time off limit IDs to associate with this time off plan.
    TimeOffLimitIds Setwrapperstring `json:"timeOffLimitIds"`


    // AutoApprovalRule - Auto approval rule for the time off plan.
    AutoApprovalRule string `json:"autoApprovalRule"`


    // DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
    DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`


    // Active - Whether this time off plan should be used by agents.
    Active bool `json:"active"`


    // Metadata - Version metadata for the time off plan
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatetimeoffplanrequest) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatetimeoffplanrequest

    if UpdatetimeoffplanrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatetimeoffplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ActivityCodeIds Setwrapperstring `json:"activityCodeIds"`
        
        TimeOffLimitIds Setwrapperstring `json:"timeOffLimitIds"`
        
        AutoApprovalRule string `json:"autoApprovalRule"`
        
        DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`
        
        Active bool `json:"active"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

