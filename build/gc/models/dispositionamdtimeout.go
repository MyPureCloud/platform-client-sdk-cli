package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DispositionamdtimeoutMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DispositionamdtimeoutDud struct { 
    


    


    


    

}

// Dispositionamdtimeout
type Dispositionamdtimeout struct { 
    // TimeoutMs - Configured AMD timeout value.
    TimeoutMs int `json:"timeoutMs"`


    // TimerStartEvent - Configured option for when to start the AMD timer, such as on line connect or speech start.
    TimerStartEvent string `json:"timerStartEvent"`


    // TimerStartTime - Timer start time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TimerStartTime time.Time `json:"timerStartTime"`


    // TimerEndTime - Timer end time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TimerEndTime time.Time `json:"timerEndTime"`

}

// String returns a JSON representation of the model
func (o *Dispositionamdtimeout) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dispositionamdtimeout) MarshalJSON() ([]byte, error) {
    type Alias Dispositionamdtimeout

    if DispositionamdtimeoutMarshalled {
        return []byte("{}"), nil
    }
    DispositionamdtimeoutMarshalled = true

    return json.Marshal(&struct {
        
        TimeoutMs int `json:"timeoutMs"`
        
        TimerStartEvent string `json:"timerStartEvent"`
        
        TimerStartTime time.Time `json:"timerStartTime"`
        
        TimerEndTime time.Time `json:"timerEndTime"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

