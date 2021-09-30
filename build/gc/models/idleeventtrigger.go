package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdleeventtriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdleeventtriggerDud struct { 
    


    

}

// Idleeventtrigger - Details about an idle event trigger
type Idleeventtrigger struct { 
    // EventName - Name of event triggered after period of inactivity.
    EventName string `json:"eventName"`


    // IdleAfterSeconds - Number of seconds of inactivity before an event is triggered.
    IdleAfterSeconds int `json:"idleAfterSeconds"`

}

// String returns a JSON representation of the model
func (o *Idleeventtrigger) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Idleeventtrigger) MarshalJSON() ([]byte, error) {
    type Alias Idleeventtrigger

    if IdleeventtriggerMarshalled {
        return []byte("{}"), nil
    }
    IdleeventtriggerMarshalled = true

    return json.Marshal(&struct { 
        EventName string `json:"eventName"`
        
        IdleAfterSeconds int `json:"idleAfterSeconds"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

