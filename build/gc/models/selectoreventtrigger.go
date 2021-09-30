package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelectoreventtriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelectoreventtriggerDud struct { 
    


    

}

// Selectoreventtrigger - Details about a selector event trigger
type Selectoreventtrigger struct { 
    // Selector - Element that triggers event.
    Selector string `json:"selector"`


    // EventName - Name of event triggered when element matching selector is interacted with.
    EventName string `json:"eventName"`

}

// String returns a JSON representation of the model
func (o *Selectoreventtrigger) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selectoreventtrigger) MarshalJSON() ([]byte, error) {
    type Alias Selectoreventtrigger

    if SelectoreventtriggerMarshalled {
        return []byte("{}"), nil
    }
    SelectoreventtriggerMarshalled = true

    return json.Marshal(&struct { 
        Selector string `json:"selector"`
        
        EventName string `json:"eventName"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

