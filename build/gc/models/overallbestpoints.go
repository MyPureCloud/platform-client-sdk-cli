package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OverallbestpointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OverallbestpointsDud struct { 
    Division Division `json:"division"`


    BestPoints []Overallbestpointsitem `json:"bestPoints"`


    PerformanceProfile Addressableentityref `json:"performanceProfile"`

}

// Overallbestpoints
type Overallbestpoints struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Overallbestpoints) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Overallbestpoints) MarshalJSON() ([]byte, error) {
    type Alias Overallbestpoints

    if OverallbestpointsMarshalled {
        return []byte("{}"), nil
    }
    OverallbestpointsMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

