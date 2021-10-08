package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SingleworkdayaveragepointsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SingleworkdayaveragepointsDud struct { 
    DateWorkday time.Time `json:"dateWorkday"`


    Division Division `json:"division"`


    AveragePoints float64 `json:"averagePoints"`


    PerformanceProfile Addressableentityref `json:"performanceProfile"`

}

// Singleworkdayaveragepoints
type Singleworkdayaveragepoints struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragepoints) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Singleworkdayaveragepoints) MarshalJSON() ([]byte, error) {
    type Alias Singleworkdayaveragepoints

    if SingleworkdayaveragepointsMarshalled {
        return []byte("{}"), nil
    }
    SingleworkdayaveragepointsMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

