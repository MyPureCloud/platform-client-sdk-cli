package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomepercentileconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomepercentileconditionDud struct { 
    


    


    

}

// Outcomepercentilecondition
type Outcomepercentilecondition struct { 
    // OutcomeId - The outcome ID.
    OutcomeId string `json:"outcomeId"`


    // MaximumPercentile - Percentile value for the selected outcome, at or above which the action map will trigger.
    MaximumPercentile float32 `json:"maximumPercentile"`


    // FallbackPercentile - Additional percentile condition, where if set, the action map will trigger if the current outcome percentile is lower or equal to the value.
    FallbackPercentile float32 `json:"fallbackPercentile"`

}

// String returns a JSON representation of the model
func (o *Outcomepercentilecondition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomepercentilecondition) MarshalJSON() ([]byte, error) {
    type Alias Outcomepercentilecondition

    if OutcomepercentileconditionMarshalled {
        return []byte("{}"), nil
    }
    OutcomepercentileconditionMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeId string `json:"outcomeId"`
        
        MaximumPercentile float32 `json:"maximumPercentile"`
        
        FallbackPercentile float32 `json:"fallbackPercentile"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

