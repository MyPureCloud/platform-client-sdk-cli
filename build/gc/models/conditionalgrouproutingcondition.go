package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgrouproutingconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgrouproutingconditionDud struct { 
    


    


    


    

}

// Conditionalgrouproutingcondition
type Conditionalgrouproutingcondition struct { 
    // Queue - The queue being evaluated for this rule.  If null, the current queue will be used.
    Queue Domainentityref `json:"queue"`


    // Metric - The queue metric being evaluated
    Metric string `json:"metric"`


    // Operator - The operator that compares the actual value against the condition value
    Operator string `json:"operator"`


    // Value - The limit value, beyond which a rule evaluates as true
    Value float64 `json:"value"`

}

// String returns a JSON representation of the model
func (o *Conditionalgrouproutingcondition) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgrouproutingcondition) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgrouproutingcondition

    if ConditionalgrouproutingconditionMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgrouproutingconditionMarshalled = true

    return json.Marshal(&struct {
        
        Queue Domainentityref `json:"queue"`
        
        Metric string `json:"metric"`
        
        Operator string `json:"operator"`
        
        Value float64 `json:"value"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

