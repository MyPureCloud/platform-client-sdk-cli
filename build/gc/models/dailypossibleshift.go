package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DailypossibleshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DailypossibleshiftDud struct { 
    


    


    


    


    


    

}

// Dailypossibleshift
type Dailypossibleshift struct { 
    // DayOfWeek - Day of the shift
    DayOfWeek string `json:"dayOfWeek"`


    // EarliestShiftStartMinutesFromMidnight - Minutes of the earliest shift start from midnight. Note that midnight is 12:00 am in the time zone specified in the timeZone field (in the top level of the response)
    EarliestShiftStartMinutesFromMidnight int `json:"earliestShiftStartMinutesFromMidnight"`


    // Required - Whether this is a required shift
    Required bool `json:"required"`


    // MinimumPaidTimeMinutes - Minimum paid time in minutes of this daily shift
    MinimumPaidTimeMinutes int `json:"minimumPaidTimeMinutes"`


    // MaximumPaidTimeMinutes - Maximum paid time in minutes of this daily shift
    MaximumPaidTimeMinutes int `json:"maximumPaidTimeMinutes"`


    // IntervalScheduleProbabilities - The percentage of being scheduled in each interval between the earliest shift start and latest shift end. Range of the values: [0, 100].
    IntervalScheduleProbabilities []int `json:"intervalScheduleProbabilities"`

}

// String returns a JSON representation of the model
func (o *Dailypossibleshift) String() string {
    
    
    
    
    
     o.IntervalScheduleProbabilities = []int{0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dailypossibleshift) MarshalJSON() ([]byte, error) {
    type Alias Dailypossibleshift

    if DailypossibleshiftMarshalled {
        return []byte("{}"), nil
    }
    DailypossibleshiftMarshalled = true

    return json.Marshal(&struct {
        
        DayOfWeek string `json:"dayOfWeek"`
        
        EarliestShiftStartMinutesFromMidnight int `json:"earliestShiftStartMinutesFromMidnight"`
        
        Required bool `json:"required"`
        
        MinimumPaidTimeMinutes int `json:"minimumPaidTimeMinutes"`
        
        MaximumPaidTimeMinutes int `json:"maximumPaidTimeMinutes"`
        
        IntervalScheduleProbabilities []int `json:"intervalScheduleProbabilities"`
        *Alias
    }{

        


        


        


        


        


        
        IntervalScheduleProbabilities: []int{0},
        

        Alias: (*Alias)(u),
    })
}

