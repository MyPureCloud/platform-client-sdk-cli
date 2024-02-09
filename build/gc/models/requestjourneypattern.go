package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RequestjourneypatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RequestjourneypatternDud struct { 
    


    


    


    


    

}

// Requestjourneypattern
type Requestjourneypattern struct { 
    // Criteria - A list of one or more criteria to satisfy.
    Criteria []Requestcriteria `json:"criteria"`


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
func (o *Requestjourneypattern) String() string {
     o.Criteria = []Requestcriteria{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Requestjourneypattern) MarshalJSON() ([]byte, error) {
    type Alias Requestjourneypattern

    if RequestjourneypatternMarshalled {
        return []byte("{}"), nil
    }
    RequestjourneypatternMarshalled = true

    return json.Marshal(&struct {
        
        Criteria []Requestcriteria `json:"criteria"`
        
        Count int `json:"count"`
        
        StreamType string `json:"streamType"`
        
        SessionType string `json:"sessionType"`
        
        EventName string `json:"eventName"`
        *Alias
    }{

        
        Criteria: []Requestcriteria{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

