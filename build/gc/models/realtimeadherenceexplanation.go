package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RealtimeadherenceexplanationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RealtimeadherenceexplanationDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Realtimeadherenceexplanation
type Realtimeadherenceexplanation struct { 
    


    // StartDate - The start timestamp of the adherence explanation in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence explanation in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Status - The status of the adherence explanation
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Realtimeadherenceexplanation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Realtimeadherenceexplanation) MarshalJSON() ([]byte, error) {
    type Alias Realtimeadherenceexplanation

    if RealtimeadherenceexplanationMarshalled {
        return []byte("{}"), nil
    }
    RealtimeadherenceexplanationMarshalled = true

    return json.Marshal(&struct { 
        
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Status string `json:"status"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

