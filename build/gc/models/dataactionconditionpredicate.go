package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataactionconditionpredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataactionconditionpredicateDud struct { 
    


    


    


    


    

}

// Dataactionconditionpredicate
type Dataactionconditionpredicate struct { 
    // OutputField - The name of an output field from the data action's output to use for this condition
    OutputField string `json:"outputField"`


    // OutputOperator - The operation with which to evaluate this condition
    OutputOperator string `json:"outputOperator"`


    // ComparisonValue - The value to compare against for this condition
    ComparisonValue string `json:"comparisonValue"`


    // Inverted - If true, inverts the result of evaluating this Predicate. Default is false.
    Inverted bool `json:"inverted"`


    // OutputFieldMissingResolution - The result of this predicate if the requested output field is missing from the data action's result
    OutputFieldMissingResolution bool `json:"outputFieldMissingResolution"`

}

// String returns a JSON representation of the model
func (o *Dataactionconditionpredicate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataactionconditionpredicate) MarshalJSON() ([]byte, error) {
    type Alias Dataactionconditionpredicate

    if DataactionconditionpredicateMarshalled {
        return []byte("{}"), nil
    }
    DataactionconditionpredicateMarshalled = true

    return json.Marshal(&struct {
        
        OutputField string `json:"outputField"`
        
        OutputOperator string `json:"outputOperator"`
        
        ComparisonValue string `json:"comparisonValue"`
        
        Inverted bool `json:"inverted"`
        
        OutputFieldMissingResolution bool `json:"outputFieldMissingResolution"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

