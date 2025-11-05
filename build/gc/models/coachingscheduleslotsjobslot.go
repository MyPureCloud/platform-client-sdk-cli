package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingscheduleslotsjobslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingscheduleslotsjobslotDud struct { 
    


    

}

// Coachingscheduleslotsjobslot
type Coachingscheduleslotsjobslot struct { 
    // DateStart - Start date and time of the slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // Schedule - The schedule information of the slot
    Schedule Coachingscheduleslotsjobschedule `json:"schedule"`

}

// String returns a JSON representation of the model
func (o *Coachingscheduleslotsjobslot) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingscheduleslotsjobslot) MarshalJSON() ([]byte, error) {
    type Alias Coachingscheduleslotsjobslot

    if CoachingscheduleslotsjobslotMarshalled {
        return []byte("{}"), nil
    }
    CoachingscheduleslotsjobslotMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        Schedule Coachingscheduleslotsjobschedule `json:"schedule"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

