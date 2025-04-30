package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgroupactivationconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgroupactivationconditionDud struct { 
    


    


    

}

// Conditionalgroupactivationcondition
type Conditionalgroupactivationcondition struct { 
    // SimpleMetric - Instructs this condition to evaluate a simple queue-level metric
    SimpleMetric Conditionalgroupactivationsimplemetric `json:"simpleMetric"`


    // Operator - The operator used to compare the actual value against the threshold value
    Operator string `json:"operator"`


    // Value - The threshold value, beyond which a rule evaluates as true
    Value float64 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Conditionalgroupactivationcondition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgroupactivationcondition) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgroupactivationcondition

    if ConditionalgroupactivationconditionMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgroupactivationconditionMarshalled = true

    return json.Marshal(&struct {
        
        SimpleMetric Conditionalgroupactivationsimplemetric `json:"simpleMetric"`
        
        Operator string `json:"operator"`
        
        Value float64 `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

