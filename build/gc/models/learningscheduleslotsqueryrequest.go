package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscheduleslotsqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscheduleslotsqueryrequestDud struct { 
    


    


    


    

}

// Learningscheduleslotsqueryrequest
type Learningscheduleslotsqueryrequest struct { 
    // Interval - Range of time to get slots for scheduling learning activities. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // LengthInMinutes - The duration of coaching appointment to schedule in 15 minutes granularity
    LengthInMinutes int `json:"lengthInMinutes"`


    // UserIds - The user IDs for which to fetch schedules. Must be only 1.
    UserIds []string `json:"userIds"`


    // InterruptibleAssignmentId - Assignment ID to exclude from consideration when determining blocked slots
    InterruptibleAssignmentId string `json:"interruptibleAssignmentId"`

}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsqueryrequest) String() string {
    
    
     o.UserIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscheduleslotsqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningscheduleslotsqueryrequest

    if LearningscheduleslotsqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningscheduleslotsqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        UserIds []string `json:"userIds"`
        
        InterruptibleAssignmentId string `json:"interruptibleAssignmentId"`
        *Alias
    }{

        


        


        
        UserIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

