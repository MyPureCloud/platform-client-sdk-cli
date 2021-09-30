package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeprobabilityconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeprobabilityconditionDud struct { 
    


    


    

}

// Outcomeprobabilitycondition
type Outcomeprobabilitycondition struct { 
    // OutcomeId - The outcome ID.
    OutcomeId string `json:"outcomeId"`


    // MaximumProbability - Probability value for the selected outcome at or above which the action map will trigger.
    MaximumProbability float32 `json:"maximumProbability"`


    // Probability - Additional probability condition, where if set, the action map will trigger if the current outcome probability is lower or equal to the value.
    Probability float32 `json:"probability"`

}

// String returns a JSON representation of the model
func (o *Outcomeprobabilitycondition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeprobabilitycondition) MarshalJSON() ([]byte, error) {
    type Alias Outcomeprobabilitycondition

    if OutcomeprobabilityconditionMarshalled {
        return []byte("{}"), nil
    }
    OutcomeprobabilityconditionMarshalled = true

    return json.Marshal(&struct { 
        OutcomeId string `json:"outcomeId"`
        
        MaximumProbability float32 `json:"maximumProbability"`
        
        Probability float32 `json:"probability"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

