package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneypatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneypatternDud struct { 
    


    


    


    


    

}

// Journeypattern
type Journeypattern struct { 
    // Criteria - A list of one or more criteria to satisfy.
    Criteria []Criteria `json:"criteria"`


    // Count - The number of times the pattern must match.
    Count int `json:"count"`


    // StreamType - The stream type for which this pattern can be matched on.
    StreamType string `json:"streamType"`


    // SessionType - The session type for which this pattern can be matched on.
    SessionType string `json:"sessionType"`


    // EventName - The name of the event for which this pattern can be matched on.
    EventName string `json:"eventName"`

}

// String returns a JSON representation of the model
func (o *Journeypattern) String() string {
    
    
     o.Criteria = []Criteria{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeypattern) MarshalJSON() ([]byte, error) {
    type Alias Journeypattern

    if JourneypatternMarshalled {
        return []byte("{}"), nil
    }
    JourneypatternMarshalled = true

    return json.Marshal(&struct { 
        Criteria []Criteria `json:"criteria"`
        
        Count int `json:"count"`
        
        StreamType string `json:"streamType"`
        
        SessionType string `json:"sessionType"`
        
        EventName string `json:"eventName"`
        
        *Alias
    }{
        

        
        Criteria: []Criteria{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

