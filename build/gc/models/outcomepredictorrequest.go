package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomepredictorrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomepredictorrequestDud struct { 
    

}

// Outcomepredictorrequest
type Outcomepredictorrequest struct { 
    // Outcome - The outcome for which this predictor will provide predictions.
    Outcome Outcomerefrequest `json:"outcome"`

}

// String returns a JSON representation of the model
func (o *Outcomepredictorrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomepredictorrequest) MarshalJSON() ([]byte, error) {
    type Alias Outcomepredictorrequest

    if OutcomepredictorrequestMarshalled {
        return []byte("{}"), nil
    }
    OutcomepredictorrequestMarshalled = true

    return json.Marshal(&struct {
        
        Outcome Outcomerefrequest `json:"outcome"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

