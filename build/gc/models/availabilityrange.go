package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailabilityrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailabilityrangeDud struct { 
    


    

}

// Availabilityrange
type Availabilityrange struct { 
    // EarliestStartMinutesFromMidnight - The earliest time of day the activity can be scheduled to begin, in minutes from midnight in the configured time zone of the business unit
    EarliestStartMinutesFromMidnight int `json:"earliestStartMinutesFromMidnight"`


    // LatestEndMinutesFromMidnight - The latest time of day the activity can be scheduled to end, in minutes from midnight in the configured time zone of the business unit
    LatestEndMinutesFromMidnight int `json:"latestEndMinutesFromMidnight"`

}

// String returns a JSON representation of the model
func (o *Availabilityrange) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availabilityrange) MarshalJSON() ([]byte, error) {
    type Alias Availabilityrange

    if AvailabilityrangeMarshalled {
        return []byte("{}"), nil
    }
    AvailabilityrangeMarshalled = true

    return json.Marshal(&struct {
        
        EarliestStartMinutesFromMidnight int `json:"earliestStartMinutesFromMidnight"`
        
        LatestEndMinutesFromMidnight int `json:"latestEndMinutesFromMidnight"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

