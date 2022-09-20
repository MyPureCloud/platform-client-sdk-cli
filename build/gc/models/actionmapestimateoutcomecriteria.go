package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapestimateoutcomecriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapestimateoutcomecriteriaDud struct { 
    


    


    

}

// Actionmapestimateoutcomecriteria
type Actionmapestimateoutcomecriteria struct { 
    // OutcomeId - ID of outcome.
    OutcomeId string `json:"outcomeId"`


    // MaxProbability - Probability value for the selected outcome at or above which the action map will trigger.
    MaxProbability float32 `json:"maxProbability"`


    // Probability - Additional probability condition, where if set, the action map will trigger if the current outcome probability is lower or equal to the value.
    Probability float32 `json:"probability"`

}

// String returns a JSON representation of the model
func (o *Actionmapestimateoutcomecriteria) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapestimateoutcomecriteria) MarshalJSON() ([]byte, error) {
    type Alias Actionmapestimateoutcomecriteria

    if ActionmapestimateoutcomecriteriaMarshalled {
        return []byte("{}"), nil
    }
    ActionmapestimateoutcomecriteriaMarshalled = true

    return json.Marshal(&struct {
        
        OutcomeId string `json:"outcomeId"`
        
        MaxProbability float32 `json:"maxProbability"`
        
        Probability float32 `json:"probability"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

