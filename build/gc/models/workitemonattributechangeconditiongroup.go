package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangeconditiongroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangeconditiongroupDud struct { 
    


    

}

// Workitemonattributechangeconditiongroup
type Workitemonattributechangeconditiongroup struct { 
    // Operator - How the conditions in this group are combined.
    Operator string `json:"operator"`


    // Conditions - The conditions to evaluate. Between 1 and 5 conditions are permitted per group.
    Conditions []Workitemonattributechangecondition `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangeconditiongroup) String() string {
    
     o.Conditions = []Workitemonattributechangecondition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangeconditiongroup) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangeconditiongroup

    if WorkitemonattributechangeconditiongroupMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangeconditiongroupMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Conditions []Workitemonattributechangecondition `json:"conditions"`
        *Alias
    }{

        


        
        Conditions: []Workitemonattributechangecondition{{}},
        

        Alias: (*Alias)(u),
    })
}

