package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediastatisticalsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediastatisticalsummaryDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Socialmediastatisticalsummary
type Socialmediastatisticalsummary struct { 
    // Max
    Max float32 `json:"max"`


    // Min
    Min float32 `json:"min"`


    // Count
    Count int `json:"count"`


    // CountNegative
    CountNegative int `json:"countNegative"`


    // CountPositive
    CountPositive int `json:"countPositive"`


    // CountNeutral
    CountNeutral int `json:"countNeutral"`


    // CountUnknown
    CountUnknown int `json:"countUnknown"`


    // Sum
    Sum float32 `json:"sum"`


    // Average
    Average float32 `json:"average"`


    // Current
    Current float32 `json:"current"`


    // Ratio
    Ratio float32 `json:"ratio"`


    // Numerator
    Numerator float32 `json:"numerator"`


    // Denominator
    Denominator float32 `json:"denominator"`


    // Target
    Target float32 `json:"target"`


    // P95
    P95 int `json:"p95"`


    // P99
    P99 int `json:"p99"`

}

// String returns a JSON representation of the model
func (o *Socialmediastatisticalsummary) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediastatisticalsummary) MarshalJSON() ([]byte, error) {
    type Alias Socialmediastatisticalsummary

    if SocialmediastatisticalsummaryMarshalled {
        return []byte("{}"), nil
    }
    SocialmediastatisticalsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Max float32 `json:"max"`
        
        Min float32 `json:"min"`
        
        Count int `json:"count"`
        
        CountNegative int `json:"countNegative"`
        
        CountPositive int `json:"countPositive"`
        
        CountNeutral int `json:"countNeutral"`
        
        CountUnknown int `json:"countUnknown"`
        
        Sum float32 `json:"sum"`
        
        Average float32 `json:"average"`
        
        Current float32 `json:"current"`
        
        Ratio float32 `json:"ratio"`
        
        Numerator float32 `json:"numerator"`
        
        Denominator float32 `json:"denominator"`
        
        Target float32 `json:"target"`
        
        P95 int `json:"p95"`
        
        P99 int `json:"p99"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

