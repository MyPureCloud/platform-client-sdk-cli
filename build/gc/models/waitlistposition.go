package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WaitlistpositionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WaitlistpositionDud struct { 
    


    


    


    

}

// Waitlistposition
type Waitlistposition struct { 
    // TimeOffRequest - The time off request for this wait list position
    TimeOffRequest Timeoffrequestreference `json:"timeOffRequest"`


    // TimeOffLimit - The time off limit for which time off request is waitlisted
    TimeOffLimit Timeofflimitreference `json:"timeOffLimit"`


    // Date - The date to which this wait list position applies, as defined by the time zone of the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Date time.Time `json:"date"`


    // WaitlistPosition - The time off request's position in the waitlist on the date. 1 means time off is the first in the waitlist
    WaitlistPosition int `json:"waitlistPosition"`

}

// String returns a JSON representation of the model
func (o *Waitlistposition) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Waitlistposition) MarshalJSON() ([]byte, error) {
    type Alias Waitlistposition

    if WaitlistpositionMarshalled {
        return []byte("{}"), nil
    }
    WaitlistpositionMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequest Timeoffrequestreference `json:"timeOffRequest"`
        
        TimeOffLimit Timeofflimitreference `json:"timeOffLimit"`
        
        Date time.Time `json:"date"`
        
        WaitlistPosition int `json:"waitlistPosition"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

