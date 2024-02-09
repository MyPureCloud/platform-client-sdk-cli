package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchjourneypatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchjourneypatternDud struct { 
    


    


    


    


    

}

// Patchjourneypattern
type Patchjourneypattern struct { 
    // Criteria - A list of one or more criteria to satisfy.
    Criteria []Patchcriteria `json:"criteria"`


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
func (o *Patchjourneypattern) String() string {
     o.Criteria = []Patchcriteria{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchjourneypattern) MarshalJSON() ([]byte, error) {
    type Alias Patchjourneypattern

    if PatchjourneypatternMarshalled {
        return []byte("{}"), nil
    }
    PatchjourneypatternMarshalled = true

    return json.Marshal(&struct {
        
        Criteria []Patchcriteria `json:"criteria"`
        
        Count int `json:"count"`
        
        StreamType string `json:"streamType"`
        
        SessionType string `json:"sessionType"`
        
        EventName string `json:"eventName"`
        *Alias
    }{

        
        Criteria: []Patchcriteria{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

