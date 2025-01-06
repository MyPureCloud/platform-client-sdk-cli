package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventdefinitionaggregatesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventdefinitionaggregatesDud struct { 
    


    

}

// Eventdefinitionaggregates
type Eventdefinitionaggregates struct { 
    // EventDefinition - Event definition.
    EventDefinition Addressableentityref `json:"eventDefinition"`


    // EventCount - How many events have occurred.
    EventCount int `json:"eventCount"`

}

// String returns a JSON representation of the model
func (o *Eventdefinitionaggregates) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventdefinitionaggregates) MarshalJSON() ([]byte, error) {
    type Alias Eventdefinitionaggregates

    if EventdefinitionaggregatesMarshalled {
        return []byte("{}"), nil
    }
    EventdefinitionaggregatesMarshalled = true

    return json.Marshal(&struct {
        
        EventDefinition Addressableentityref `json:"eventDefinition"`
        
        EventCount int `json:"eventCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

