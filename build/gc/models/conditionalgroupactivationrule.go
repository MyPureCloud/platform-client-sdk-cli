package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgroupactivationruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgroupactivationruleDud struct { 
    


    


    

}

// Conditionalgroupactivationrule
type Conditionalgroupactivationrule struct { 
    // Conditions - The list of conditions used in this rule
    Conditions []Conditionalgroupactivationcondition `json:"conditions"`


    // ConditionExpression - A string expression that defines the relationships of conditions in this rule
    ConditionExpression string `json:"conditionExpression"`


    // Groups - The group(s) that this rule activates (if rule evaluates as true) or deactivates (if rule evaluates as false)
    Groups []Membergroup `json:"groups"`

}

// String returns a JSON representation of the model
func (o *Conditionalgroupactivationrule) String() string {
     o.Conditions = []Conditionalgroupactivationcondition{{}} 
    
     o.Groups = []Membergroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgroupactivationrule) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgroupactivationrule

    if ConditionalgroupactivationruleMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgroupactivationruleMarshalled = true

    return json.Marshal(&struct {
        
        Conditions []Conditionalgroupactivationcondition `json:"conditions"`
        
        ConditionExpression string `json:"conditionExpression"`
        
        Groups []Membergroup `json:"groups"`
        *Alias
    }{

        
        Conditions: []Conditionalgroupactivationcondition{{}},
        


        


        
        Groups: []Membergroup{{}},
        

        Alias: (*Alias)(u),
    })
}

