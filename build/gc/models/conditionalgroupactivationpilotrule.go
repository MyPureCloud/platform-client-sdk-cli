package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgroupactivationpilotruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgroupactivationpilotruleDud struct { 
    


    

}

// Conditionalgroupactivationpilotrule
type Conditionalgroupactivationpilotrule struct { 
    // Conditions - The list of conditions used in this rule
    Conditions []Conditionalgroupactivationcondition `json:"conditions"`


    // ConditionExpression - A string expression that defines the relationships of conditions in this rule
    ConditionExpression string `json:"conditionExpression"`

}

// String returns a JSON representation of the model
func (o *Conditionalgroupactivationpilotrule) String() string {
     o.Conditions = []Conditionalgroupactivationcondition{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgroupactivationpilotrule) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgroupactivationpilotrule

    if ConditionalgroupactivationpilotruleMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgroupactivationpilotruleMarshalled = true

    return json.Marshal(&struct {
        
        Conditions []Conditionalgroupactivationcondition `json:"conditions"`
        
        ConditionExpression string `json:"conditionExpression"`
        *Alias
    }{

        
        Conditions: []Conditionalgroupactivationcondition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

