package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeofflimitvaluesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeofflimitvaluesDud struct { 
    


    


    


    

}

// Timeofflimitvalues
type Timeofflimitvalues struct { 
    // LimitMinutes - Time-off limit values in minutes per granularity interval
    LimitMinutes []int `json:"limitMinutes"`


    // AllocatedMinutes - Allocated time-off minutes per granularity interval
    AllocatedMinutes []int `json:"allocatedMinutes"`


    // WaitlistedMinutes - Waitlisted time-off minutes per granularity interval
    WaitlistedMinutes []int `json:"waitlistedMinutes"`


    // WaitlistedRequests - The current number of waitlisted time-off requests per granularity interval
    WaitlistedRequests []int `json:"waitlistedRequests"`

}

// String returns a JSON representation of the model
func (o *Timeofflimitvalues) String() string {
     o.LimitMinutes = []int{0} 
     o.AllocatedMinutes = []int{0} 
     o.WaitlistedMinutes = []int{0} 
     o.WaitlistedRequests = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeofflimitvalues) MarshalJSON() ([]byte, error) {
    type Alias Timeofflimitvalues

    if TimeofflimitvaluesMarshalled {
        return []byte("{}"), nil
    }
    TimeofflimitvaluesMarshalled = true

    return json.Marshal(&struct {
        
        LimitMinutes []int `json:"limitMinutes"`
        
        AllocatedMinutes []int `json:"allocatedMinutes"`
        
        WaitlistedMinutes []int `json:"waitlistedMinutes"`
        
        WaitlistedRequests []int `json:"waitlistedRequests"`
        *Alias
    }{

        
        LimitMinutes: []int{0},
        


        
        AllocatedMinutes: []int{0},
        


        
        WaitlistedMinutes: []int{0},
        


        
        WaitlistedRequests: []int{0},
        

        Alias: (*Alias)(u),
    })
}

