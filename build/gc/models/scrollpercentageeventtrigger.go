package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScrollpercentageeventtriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScrollpercentageeventtriggerDud struct { 
    


    

}

// Scrollpercentageeventtrigger - Details about a scroll percentage event trigger
type Scrollpercentageeventtrigger struct { 
    // Percentage - Percentage of a webpage at which an event is triggered.
    Percentage int `json:"percentage"`


    // EventName - Name of event triggered after scrolling to the specified percentage.
    EventName string `json:"eventName"`

}

// String returns a JSON representation of the model
func (o *Scrollpercentageeventtrigger) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scrollpercentageeventtrigger) MarshalJSON() ([]byte, error) {
    type Alias Scrollpercentageeventtrigger

    if ScrollpercentageeventtriggerMarshalled {
        return []byte("{}"), nil
    }
    ScrollpercentageeventtriggerMarshalled = true

    return json.Marshal(&struct {
        
        Percentage int `json:"percentage"`
        
        EventName string `json:"eventName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

