package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangeconditionupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangeconditionupdateDud struct { 
    


    


    

}

// Workitemonattributechangeconditionupdate
type Workitemonattributechangeconditionupdate struct { 
    // Attribute - The name of the workitem attribute whose change will be evaluated as part of the rule.
    Attribute string `json:"attribute"`


    // NewValue - The new value of the attribute. If the attribute is updated to this value this part of the condition will be met.
    NewValue string `json:"newValue"`


    // OldValue - The old value of the attribute. If the attribute was updated from this value this part of the condition will be met.
    OldValue string `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangeconditionupdate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangeconditionupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangeconditionupdate

    if WorkitemonattributechangeconditionupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangeconditionupdateMarshalled = true

    return json.Marshal(&struct {
        
        Attribute string `json:"attribute"`
        
        NewValue string `json:"newValue"`
        
        OldValue string `json:"oldValue"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

