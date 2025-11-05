package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscheduleslotsjobscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscheduleslotsjobscheduleDud struct { 
    


    


    


    

}

// Learningscheduleslotsjobschedule
type Learningscheduleslotsjobschedule struct { 
    // Id
    Id string `json:"id"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    WeekDate time.Time `json:"weekDate"`


    // WeekCount - The number of weeks in this schedule
    WeekCount int `json:"weekCount"`

}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsjobschedule) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscheduleslotsjobschedule) MarshalJSON() ([]byte, error) {
    type Alias Learningscheduleslotsjobschedule

    if LearningscheduleslotsjobscheduleMarshalled {
        return []byte("{}"), nil
    }
    LearningscheduleslotsjobscheduleMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        WeekDate time.Time `json:"weekDate"`
        
        WeekCount int `json:"weekCount"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

