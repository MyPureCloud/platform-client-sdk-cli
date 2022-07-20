package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgrouproutingruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgrouproutingruleDud struct { 
    


    


    


    


    


    

}

// Conditionalgrouproutingrule
type Conditionalgrouproutingrule struct { 
    // Queue - The queue being evaluated for this rule.  For rule 1, this is always the current queue, so should not be specified.
    Queue Domainentityref `json:"queue"`


    // Metric - The queue metric being evaluated
    Metric string `json:"metric"`


    // Operator - The operator that compares the actual value against the condition value
    Operator string `json:"operator"`


    // ConditionValue - The limit value, beyond which a rule evaluates as true
    ConditionValue float64 `json:"conditionValue"`


    // Groups - The group(s) to activate if the rule evaluates as true
    Groups []Membergroup `json:"groups"`


    // WaitSeconds - The number of seconds to wait in this rule, if it evaluates as true, before evaluating the next rule.  For the final rule, this is ignored, so need not be specified.
    WaitSeconds int `json:"waitSeconds"`

}

// String returns a JSON representation of the model
func (o *Conditionalgrouproutingrule) String() string {
    
    
    
    
     o.Groups = []Membergroup{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgrouproutingrule) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgrouproutingrule

    if ConditionalgrouproutingruleMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgrouproutingruleMarshalled = true

    return json.Marshal(&struct {
        
        Queue Domainentityref `json:"queue"`
        
        Metric string `json:"metric"`
        
        Operator string `json:"operator"`
        
        ConditionValue float64 `json:"conditionValue"`
        
        Groups []Membergroup `json:"groups"`
        
        WaitSeconds int `json:"waitSeconds"`
        *Alias
    }{

        


        


        


        


        
        Groups: []Membergroup{{}},
        


        

        Alias: (*Alias)(u),
    })
}

