package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DispositionsilentcalltimeoutMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DispositionsilentcalltimeoutDud struct { 
    


    


    

}

// Dispositionsilentcalltimeout
type Dispositionsilentcalltimeout struct { 
    // TimeoutMs - Configured silent call timeout value.
    TimeoutMs int `json:"timeoutMs"`


    // TimerStartTime - Timer start time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TimerStartTime time.Time `json:"timerStartTime"`


    // TimerEndTime - Timer end time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TimerEndTime time.Time `json:"timerEndTime"`

}

// String returns a JSON representation of the model
func (o *Dispositionsilentcalltimeout) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dispositionsilentcalltimeout) MarshalJSON() ([]byte, error) {
    type Alias Dispositionsilentcalltimeout

    if DispositionsilentcalltimeoutMarshalled {
        return []byte("{}"), nil
    }
    DispositionsilentcalltimeoutMarshalled = true

    return json.Marshal(&struct {
        
        TimeoutMs int `json:"timeoutMs"`
        
        TimerStartTime time.Time `json:"timerStartTime"`
        
        TimerEndTime time.Time `json:"timerEndTime"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

