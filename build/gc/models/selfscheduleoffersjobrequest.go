package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleoffersjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleoffersjobrequestDud struct { 
    


    

}

// Selfscheduleoffersjobrequest
type Selfscheduleoffersjobrequest struct { 
    // ShiftId - The ID of the shift that the user wants to self-schedule
    ShiftId string `json:"shiftId"`


    // Schedule - A reference to the associated schedule to which this shift belongs
    Schedule Buschedulereference `json:"schedule"`

}

// String returns a JSON representation of the model
func (o *Selfscheduleoffersjobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleoffersjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleoffersjobrequest

    if SelfscheduleoffersjobrequestMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleoffersjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ShiftId string `json:"shiftId"`
        
        Schedule Buschedulereference `json:"schedule"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

