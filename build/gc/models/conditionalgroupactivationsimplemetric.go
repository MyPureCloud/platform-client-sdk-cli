package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConditionalgroupactivationsimplemetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConditionalgroupactivationsimplemetricDud struct { 
    


    

}

// Conditionalgroupactivationsimplemetric
type Conditionalgroupactivationsimplemetric struct { 
    // Metric - The queue metric being evaluated
    Metric string `json:"metric"`


    // Queue - The queue being evaluated for this rule.  If null, the current queue will be used.
    Queue Domainentityref `json:"queue"`

}

// String returns a JSON representation of the model
func (o *Conditionalgroupactivationsimplemetric) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conditionalgroupactivationsimplemetric) MarshalJSON() ([]byte, error) {
    type Alias Conditionalgroupactivationsimplemetric

    if ConditionalgroupactivationsimplemetricMarshalled {
        return []byte("{}"), nil
    }
    ConditionalgroupactivationsimplemetricMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Queue Domainentityref `json:"queue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

