package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomequantileconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomequantileconditionDud struct { 
    


    


    

}

// Outcomequantilecondition
type Outcomequantilecondition struct { 
    // OutcomeId - The outcome ID.
    OutcomeId string `json:"outcomeId"`


    // MaxQuantileThreshold - This Outcome Quantile Condition is met when sessionMaxQuantile of the OutcomeScore is above this value, (unless fallbackQuantile is set). Range 0.00-1.00
    MaxQuantileThreshold float32 `json:"maxQuantileThreshold"`


    // FallbackQuantileThreshold - (Optional) If set, this Condition is met when maxQuantileThreshold is met, AND the current quantile of the OutcomeScore is below this fallbackQuantileThreshold. Range 0.00-1.00
    FallbackQuantileThreshold float32 `json:"fallbackQuantileThreshold"`

}

// String returns a JSON representation of the model
func (o *Outcomequantilecondition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomequantilecondition) MarshalJSON() ([]byte, error) {
    type Alias Outcomequantilecondition

    if OutcomequantileconditionMarshalled {
        return []byte("{}"), nil
    }
    OutcomequantileconditionMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeId string `json:"outcomeId"`
        
        MaxQuantileThreshold float32 `json:"maxQuantileThreshold"`
        
        FallbackQuantileThreshold float32 `json:"fallbackQuantileThreshold"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

