package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabletimeoffrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabletimeoffrangeDud struct { 
    


    


    


    


    


    

}

// Availabletimeoffrange
type Availabletimeoffrange struct { 
    // TimeOffLimit - The time off limit
    TimeOffLimit Timeofflimitreference `json:"timeOffLimit"`


    // StartDate - Start date of the requested date range. The end date is determined by the size of interval list. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    StartDate time.Time `json:"startDate"`


    // Granularity - Granularity choice for time off limit
    Granularity string `json:"granularity"`


    // AvailableMinutesPerInterval - The list of available time off values in minutes per granularity interval
    AvailableMinutesPerInterval []int `json:"availableMinutesPerInterval"`


    // WaitlistedRequestsPerInterval - The current number of waitlisted time off requests for every interval per granularity
    WaitlistedRequestsPerInterval []int `json:"waitlistedRequestsPerInterval"`


    // WaitlistEnabled - Whether the time off request can be waitlisted
    WaitlistEnabled bool `json:"waitlistEnabled"`

}

// String returns a JSON representation of the model
func (o *Availabletimeoffrange) String() string {
    
    
    
     o.AvailableMinutesPerInterval = []int{0} 
     o.WaitlistedRequestsPerInterval = []int{0} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabletimeoffrange) MarshalJSON() ([]byte, error) {
    type Alias Availabletimeoffrange

    if AvailabletimeoffrangeMarshalled {
        return []byte("{}"), nil
    }
    AvailabletimeoffrangeMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffLimit Timeofflimitreference `json:"timeOffLimit"`
        
        StartDate time.Time `json:"startDate"`
        
        Granularity string `json:"granularity"`
        
        AvailableMinutesPerInterval []int `json:"availableMinutesPerInterval"`
        
        WaitlistedRequestsPerInterval []int `json:"waitlistedRequestsPerInterval"`
        
        WaitlistEnabled bool `json:"waitlistEnabled"`
        *Alias
    }{

        


        


        


        
        AvailableMinutesPerInterval: []int{0},
        


        
        WaitlistedRequestsPerInterval: []int{0},
        


        

        Alias: (*Alias)(u),
    })
}

