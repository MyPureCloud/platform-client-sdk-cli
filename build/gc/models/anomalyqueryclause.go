package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalyqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalyqueryclauseDud struct { 
    


    

}

// Anomalyqueryclause
type Anomalyqueryclause struct { 
    // Operator - The boolean operation to apply to the provided predicates
    Operator string `json:"operator"`


    // Predicates - The predicates within the query clause
    Predicates []Anomalyquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Anomalyqueryclause) String() string {
    
     o.Predicates = []Anomalyquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalyqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Anomalyqueryclause

    if AnomalyqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    AnomalyqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Predicates []Anomalyquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Anomalyquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

