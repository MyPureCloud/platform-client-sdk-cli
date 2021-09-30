package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventconditionDud struct { 
    


    


    


    


    


    

}

// Eventcondition
type Eventcondition struct { 
    // Key - The event key.
    Key string `json:"key"`


    // Values - The event values.
    Values []string `json:"values"`


    // Operator - The comparison operator.
    Operator string `json:"operator"`


    // StreamType - The stream type for which this condition can be satisfied.
    StreamType string `json:"streamType"`


    // SessionType - The session type for which this condition can be satisfied.
    SessionType string `json:"sessionType"`


    // EventName - The name of the event for which this condition can be satisfied.
    EventName string `json:"eventName"`

}

// String returns a JSON representation of the model
func (o *Eventcondition) String() string {
    
    
    
    
    
    
     o.Values = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventcondition) MarshalJSON() ([]byte, error) {
    type Alias Eventcondition

    if EventconditionMarshalled {
        return []byte("{}"), nil
    }
    EventconditionMarshalled = true

    return json.Marshal(&struct { 
        Key string `json:"key"`
        
        Values []string `json:"values"`
        
        Operator string `json:"operator"`
        
        StreamType string `json:"streamType"`
        
        SessionType string `json:"sessionType"`
        
        EventName string `json:"eventName"`
        
        *Alias
    }{
        

        

        

        
        Values: []string{""},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

