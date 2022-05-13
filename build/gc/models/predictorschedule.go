package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictorscheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictorscheduleDud struct { 
    


    DateStarted time.Time `json:"dateStarted"`

}

// Predictorschedule
type Predictorschedule struct { 
    // ScheduleType - The predictor schedule type.
    ScheduleType string `json:"scheduleType"`


    

}

// String returns a JSON representation of the model
func (o *Predictorschedule) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictorschedule) MarshalJSON() ([]byte, error) {
    type Alias Predictorschedule

    if PredictorscheduleMarshalled {
        return []byte("{}"), nil
    }
    PredictorscheduleMarshalled = true

    return json.Marshal(&struct {
        
        ScheduleType string `json:"scheduleType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

