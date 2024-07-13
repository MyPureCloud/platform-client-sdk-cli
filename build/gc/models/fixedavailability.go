package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FixedavailabilityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FixedavailabilityDud struct { 
    


    


    

}

// Fixedavailability
type Fixedavailability struct { 
    // AvailabilityRange - The range of time of day the activity can be scheduled
    AvailabilityRange Availabilityrange `json:"availabilityRange"`


    // DateRange - The range of date for which the activity plan could be scheduled
    DateRange Requiredlocaldaterange `json:"dateRange"`


    // DaysOfWeek - The days of week available for scheduling. Empty list or null means daysOfWeek is not considered
    DaysOfWeek []string `json:"daysOfWeek"`

}

// String returns a JSON representation of the model
func (o *Fixedavailability) String() string {
    
    
     o.DaysOfWeek = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fixedavailability) MarshalJSON() ([]byte, error) {
    type Alias Fixedavailability

    if FixedavailabilityMarshalled {
        return []byte("{}"), nil
    }
    FixedavailabilityMarshalled = true

    return json.Marshal(&struct {
        
        AvailabilityRange Availabilityrange `json:"availabilityRange"`
        
        DateRange Requiredlocaldaterange `json:"dateRange"`
        
        DaysOfWeek []string `json:"daysOfWeek"`
        *Alias
    }{

        


        


        
        DaysOfWeek: []string{""},
        

        Alias: (*Alias)(u),
    })
}

