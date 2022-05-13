package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomeeventscoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomeeventscoreDud struct { 
    


    


    

}

// Outcomeeventscore
type Outcomeeventscore struct { 
    // Outcome - The outcome that the score was calculated for.
    Outcome Addressableentityref `json:"outcome"`


    // SessionMaxProbability - Represents the max probability reached in the session.
    SessionMaxProbability float32 `json:"sessionMaxProbability"`


    // Probability - Represents the likelihood of a customer reaching or achieving a given outcome.
    Probability float32 `json:"probability"`

}

// String returns a JSON representation of the model
func (o *Outcomeeventscore) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomeeventscore) MarshalJSON() ([]byte, error) {
    type Alias Outcomeeventscore

    if OutcomeeventscoreMarshalled {
        return []byte("{}"), nil
    }
    OutcomeeventscoreMarshalled = true

    return json.Marshal(&struct {
        
        Outcome Addressableentityref `json:"outcome"`
        
        SessionMaxProbability float32 `json:"sessionMaxProbability"`
        
        Probability float32 `json:"probability"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

