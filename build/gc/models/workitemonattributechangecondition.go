package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangeconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangeconditionDud struct { 
    


    


    


    

}

// Workitemonattributechangecondition
type Workitemonattributechangecondition struct { 
    // Attribute - The name of the workitem attribute whose change will be evaluated as part of the rule.
    Attribute string `json:"attribute"`


    // NewValue - The value the attribute's new value is compared against. Compared for equality unless an operator is supplied. Null matches the attribute having been cleared, on clearable attributes.
    NewValue string `json:"newValue"`


    // OldValue - The value the attribute is transitioning from. Not permitted for the priority attribute. Omitted or null matches any previous value.
    OldValue string `json:"oldValue"`


    // Operator - How the attribute's new value is compared against newValue. Required for the priority attribute; not permitted otherwise.
    Operator string `json:"operator"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangecondition) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangecondition) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangecondition

    if WorkitemonattributechangeconditionMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangeconditionMarshalled = true

    return json.Marshal(&struct {
        
        Attribute string `json:"attribute"`
        
        NewValue string `json:"newValue"`
        
        OldValue string `json:"oldValue"`
        
        Operator string `json:"operator"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

