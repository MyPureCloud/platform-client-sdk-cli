package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningscheduleslotsjobresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningscheduleslotsjobresultDud struct { 
    


    


    

}

// Learningscheduleslotsjobresult
type Learningscheduleslotsjobresult struct { 
    // Interval - The interval of the job. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Status - The status of the job
    Status string `json:"status"`


    // Slot - The slot found from the job
    Slot Learningscheduleslotsjobslot `json:"slot"`

}

// String returns a JSON representation of the model
func (o *Learningscheduleslotsjobresult) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningscheduleslotsjobresult) MarshalJSON() ([]byte, error) {
    type Alias Learningscheduleslotsjobresult

    if LearningscheduleslotsjobresultMarshalled {
        return []byte("{}"), nil
    }
    LearningscheduleslotsjobresultMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Status string `json:"status"`
        
        Slot Learningscheduleslotsjobslot `json:"slot"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

