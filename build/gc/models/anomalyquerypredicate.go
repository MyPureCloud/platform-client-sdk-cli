package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalyquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalyquerypredicateDud struct { 
    


    


    

}

// Anomalyquerypredicate
type Anomalyquerypredicate struct { 
    // Dimension - The dimension on which to filter
    Dimension string `json:"dimension"`


    // Operator - The operator to use for the predicate
    Operator string `json:"operator"`


    // Value - The value against which to compare
    Value interface{} `json:"value"`

}

// String returns a JSON representation of the model
func (o *Anomalyquerypredicate) String() string {
    
    
     o.Value = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalyquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Anomalyquerypredicate

    if AnomalyquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    AnomalyquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        Dimension string `json:"dimension"`
        
        Operator string `json:"operator"`
        
        Value interface{} `json:"value"`
        *Alias
    }{

        


        


        
        Value: Interface{},
        

        Alias: (*Alias)(u),
    })
}

