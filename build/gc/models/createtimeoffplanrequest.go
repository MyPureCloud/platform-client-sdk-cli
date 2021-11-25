package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatetimeoffplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatetimeoffplanrequestDud struct { 
    


    


    


    


    


    

}

// Createtimeoffplanrequest
type Createtimeoffplanrequest struct { 
    // Name - The name of this time off plan.
    Name string `json:"name"`


    // ActivityCodeIds - The set of activity code IDs to associate with this time off plan.
    ActivityCodeIds []string `json:"activityCodeIds"`


    // TimeOffLimitIds - The set of time off limit IDs to associate with this time off plan.
    TimeOffLimitIds []string `json:"timeOffLimitIds"`


    // AutoApprovalRule - Auto approval rule for the time off plan.
    AutoApprovalRule string `json:"autoApprovalRule"`


    // DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
    DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`


    // Active - Whether this time off plan should be used by agents.
    Active bool `json:"active"`

}

// String returns a JSON representation of the model
func (o *Createtimeoffplanrequest) String() string {
    
    
    
    
    
    
     o.ActivityCodeIds = []string{""} 
    
    
    
     o.TimeOffLimitIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createtimeoffplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Createtimeoffplanrequest

    if CreatetimeoffplanrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatetimeoffplanrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        TimeOffLimitIds []string `json:"timeOffLimitIds"`
        
        AutoApprovalRule string `json:"autoApprovalRule"`
        
        DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`
        
        Active bool `json:"active"`
        
        *Alias
    }{
        

        

        

        
        ActivityCodeIds: []string{""},
        

        

        
        TimeOffLimitIds: []string{""},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

