package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DigitalruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DigitalruleDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Digitalrule
type Digitalrule struct { 
    


    // Name - The name of the rule.
    Name string `json:"name"`


    // Order - The ranked order of the rule. Rules are processed from lowest number to highest.
    Order int `json:"order"`


    // Category - The category of the rule.
    Category string `json:"category"`


    // Conditions - A list of conditions to evaluate. All of the Conditions must evaluate to true to trigger the actions.
    Conditions []Digitalcondition `json:"conditions"`


    // Actions - The list of actions to be taken if all conditions are true.
    Actions []Digitalaction `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Digitalrule) String() string {
    
    
    
     o.Conditions = []Digitalcondition{{}} 
     o.Actions = []Digitalaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Digitalrule) MarshalJSON() ([]byte, error) {
    type Alias Digitalrule

    if DigitalruleMarshalled {
        return []byte("{}"), nil
    }
    DigitalruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Order int `json:"order"`
        
        Category string `json:"category"`
        
        Conditions []Digitalcondition `json:"conditions"`
        
        Actions []Digitalaction `json:"actions"`
        *Alias
    }{

        


        


        


        


        
        Conditions: []Digitalcondition{{}},
        


        
        Actions: []Digitalaction{{}},
        

        Alias: (*Alias)(u),
    })
}

