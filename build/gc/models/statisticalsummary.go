package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StatisticalsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StatisticalsummaryDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Statisticalsummary
type Statisticalsummary struct { 
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


    // Sum
    Sum float32 `json:"sum"`


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

}

// String returns a JSON representation of the model
func (o *Statisticalsummary) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Statisticalsummary) MarshalJSON() ([]byte, error) {
    type Alias Statisticalsummary

    if StatisticalsummaryMarshalled {
        return []byte("{}"), nil
    }
    StatisticalsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Max float32 `json:"max"`
        
        Min float32 `json:"min"`
        
        Count int `json:"count"`
        
        CountNegative int `json:"countNegative"`
        
        CountPositive int `json:"countPositive"`
        
        Sum float32 `json:"sum"`
        
        Current float32 `json:"current"`
        
        Ratio float32 `json:"ratio"`
        
        Numerator float32 `json:"numerator"`
        
        Denominator float32 `json:"denominator"`
        
        Target float32 `json:"target"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

