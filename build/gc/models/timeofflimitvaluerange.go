package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeofflimitvaluerangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeofflimitvaluerangeDud struct { 
    


    


    


    


    


    


    


    

}

// Timeofflimitvaluerange - The date range with lists of time off limit, allocated and waitlisted minutes.
type Timeofflimitvaluerange struct { 
    // TimeOffLimit - The ID of the time off limit
    TimeOffLimit Timeofflimitreference `json:"timeOffLimit"`


    // StartDate - Start date of the requested date range, in ISO-8601 format. The end date is determined by the size of interval lists
    StartDate time.Time `json:"startDate"`


    // Granularity - Granularity choice for time off limit
    Granularity string `json:"granularity"`


    // LimitMinutesPerInterval - A list of time off limit values in minutes per granularity interval
    LimitMinutesPerInterval []int `json:"limitMinutesPerInterval"`


    // AllocatedMinutesPerInterval - A list of allocated time off minutes per granularity interval
    AllocatedMinutesPerInterval []int `json:"allocatedMinutesPerInterval"`


    // WaitlistedMinutesPerInterval - A list of waitlisted time off minutes per granularity interval
    WaitlistedMinutesPerInterval []int `json:"waitlistedMinutesPerInterval"`


    // WaitlistedRequestsPerInterval - The current number of waitlisted time off requests for every interval per granularity
    WaitlistedRequestsPerInterval []int `json:"waitlistedRequestsPerInterval"`


    // Metadata - Version metadata for the time off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Timeofflimitvaluerange) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.LimitMinutesPerInterval = []int{0} 
    
    
    
     o.AllocatedMinutesPerInterval = []int{0} 
    
    
    
     o.WaitlistedMinutesPerInterval = []int{0} 
    
    
    
     o.WaitlistedRequestsPerInterval = []int{0} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeofflimitvaluerange) MarshalJSON() ([]byte, error) {
    type Alias Timeofflimitvaluerange

    if TimeofflimitvaluerangeMarshalled {
        return []byte("{}"), nil
    }
    TimeofflimitvaluerangeMarshalled = true

    return json.Marshal(&struct { 
        TimeOffLimit Timeofflimitreference `json:"timeOffLimit"`
        
        StartDate time.Time `json:"startDate"`
        
        Granularity string `json:"granularity"`
        
        LimitMinutesPerInterval []int `json:"limitMinutesPerInterval"`
        
        AllocatedMinutesPerInterval []int `json:"allocatedMinutesPerInterval"`
        
        WaitlistedMinutesPerInterval []int `json:"waitlistedMinutesPerInterval"`
        
        WaitlistedRequestsPerInterval []int `json:"waitlistedRequestsPerInterval"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        LimitMinutesPerInterval: []int{0},
        

        

        
        AllocatedMinutesPerInterval: []int{0},
        

        

        
        WaitlistedMinutesPerInterval: []int{0},
        

        

        
        WaitlistedRequestsPerInterval: []int{0},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

