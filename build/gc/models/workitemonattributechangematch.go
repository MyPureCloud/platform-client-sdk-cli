package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangematchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangematchDud struct { 
    


    

}

// Workitemonattributechangematch
type Workitemonattributechangematch struct { 
    // Operator - How the results of the groups are combined. The conditions within each group are combined separately, by that group's own operator.
    Operator string `json:"operator"`


    // Groups - The groups to evaluate. Between 1 and 5 groups are permitted.
    Groups []Workitemonattributechangeconditiongroup `json:"groups"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangematch) String() string {
    
     o.Groups = []Workitemonattributechangeconditiongroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangematch) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangematch

    if WorkitemonattributechangematchMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangematchMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Groups []Workitemonattributechangeconditiongroup `json:"groups"`
        *Alias
    }{

        


        
        Groups: []Workitemonattributechangeconditiongroup{{}},
        

        Alias: (*Alias)(u),
    })
}

