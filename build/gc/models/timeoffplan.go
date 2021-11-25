package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffplanDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Timeoffplan
type Timeoffplan struct { 
    


    // Name - The name of this time off plan.
    Name string `json:"name"`


    // ActivityCodeIds - The set of activity code IDs associated with this time off plan.
    ActivityCodeIds []string `json:"activityCodeIds"`


    // TimeOffLimits - The set of time off limit IDs associated with this time off plan.
    TimeOffLimits []Timeofflimitreference `json:"timeOffLimits"`


    // AutoApprovalRule - Auto approval rule for this time off plan
    AutoApprovalRule string `json:"autoApprovalRule"`


    // DaysBeforeStartToExpireFromWaitlist - The number of days before the time off request start date for when the request will be expired from the waitlist.
    DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`


    // Active - Whether this time off plan is currently being used by agents.
    Active bool `json:"active"`


    // Metadata - Version metadata for the time off plan.
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Timeoffplan) String() string {
    
    
    
    
    
    
    
    
     o.ActivityCodeIds = []string{""} 
    
    
    
     o.TimeOffLimits = []Timeofflimitreference{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffplan) MarshalJSON() ([]byte, error) {
    type Alias Timeoffplan

    if TimeoffplanMarshalled {
        return []byte("{}"), nil
    }
    TimeoffplanMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        TimeOffLimits []Timeofflimitreference `json:"timeOffLimits"`
        
        AutoApprovalRule string `json:"autoApprovalRule"`
        
        DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`
        
        Active bool `json:"active"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        ActivityCodeIds: []string{""},
        

        

        
        TimeOffLimits: []Timeofflimitreference{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

