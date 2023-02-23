package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomepredictorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomepredictorDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Outcomepredictor
type Outcomepredictor struct { 
    


    // Outcome - The outcome for which this predictor will provide predictions.
    Outcome Outcomeref `json:"outcome"`


    

}

// String returns a JSON representation of the model
func (o *Outcomepredictor) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomepredictor) MarshalJSON() ([]byte, error) {
    type Alias Outcomepredictor

    if OutcomepredictorMarshalled {
        return []byte("{}"), nil
    }
    OutcomepredictorMarshalled = true

    return json.Marshal(&struct {
        
        Outcome Outcomeref `json:"outcome"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

