package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscheduleslotsjobslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscheduleslotsjobslotDud struct { 
    


    

}

// Learningscheduleslotsjobslot
type Learningscheduleslotsjobslot struct { 
    // DateStart - Start date and time of the slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // Schedule - The schedule information of the slot
    Schedule Learningscheduleslotsjobschedule `json:"schedule"`

}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsjobslot) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscheduleslotsjobslot) MarshalJSON() ([]byte, error) {
    type Alias Learningscheduleslotsjobslot

    if LearningscheduleslotsjobslotMarshalled {
        return []byte("{}"), nil
    }
    LearningscheduleslotsjobslotMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        Schedule Learningscheduleslotsjobschedule `json:"schedule"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

