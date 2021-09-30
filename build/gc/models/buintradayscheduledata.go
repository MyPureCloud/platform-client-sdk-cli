package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuintradayscheduledataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuintradayscheduledataDud struct { 
    

}

// Buintradayscheduledata
type Buintradayscheduledata struct { 
    // OnQueueTimeSeconds - The total on-queue time in seconds for all agents in this group
    OnQueueTimeSeconds int `json:"onQueueTimeSeconds"`

}

// String returns a JSON representation of the model
func (o *Buintradayscheduledata) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buintradayscheduledata) MarshalJSON() ([]byte, error) {
    type Alias Buintradayscheduledata

    if BuintradayscheduledataMarshalled {
        return []byte("{}"), nil
    }
    BuintradayscheduledataMarshalled = true

    return json.Marshal(&struct { 
        OnQueueTimeSeconds int `json:"onQueueTimeSeconds"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

