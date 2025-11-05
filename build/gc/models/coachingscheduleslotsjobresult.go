package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingscheduleslotsjobresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingscheduleslotsjobresultDud struct { 
    


    


    


    

}

// Coachingscheduleslotsjobresult
type Coachingscheduleslotsjobresult struct { 
    // Interval - The interval of the job. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Status - The status of the job
    Status string `json:"status"`


    // Slot - The slot found from the job
    Slot Coachingscheduleslotsjobslot `json:"slot"`


    // HasConflict - True if the slot conflicts with a Coaching Appointment
    HasConflict bool `json:"hasConflict"`

}

// String returns a JSON representation of the model
func (o *Coachingscheduleslotsjobresult) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingscheduleslotsjobresult) MarshalJSON() ([]byte, error) {
    type Alias Coachingscheduleslotsjobresult

    if CoachingscheduleslotsjobresultMarshalled {
        return []byte("{}"), nil
    }
    CoachingscheduleslotsjobresultMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Status string `json:"status"`
        
        Slot Coachingscheduleslotsjobslot `json:"slot"`
        
        HasConflict bool `json:"hasConflict"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

