package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialerruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialerruleDud struct { 
    Id string `json:"id"`


    


    


    


    


    

}

// Dialerrule
type Dialerrule struct { 
    


    // Name - The name of the rule.
    Name string `json:"name"`


    // Order - The ranked order of the rule. Rules are processed from lowest number to highest.
    Order int `json:"order"`


    // Category - The category of the rule.
    Category string `json:"category"`


    // Conditions - A list of Conditions. All of the Conditions must evaluate to true to trigger the actions.
    Conditions []Condition `json:"conditions"`


    // Actions - The list of actions to be taken if the conditions are true.
    Actions []Dialeraction `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Dialerrule) String() string {
    
    
    
     o.Conditions = []Condition{{}} 
     o.Actions = []Dialeraction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialerrule) MarshalJSON() ([]byte, error) {
    type Alias Dialerrule

    if DialerruleMarshalled {
        return []byte("{}"), nil
    }
    DialerruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Order int `json:"order"`
        
        Category string `json:"category"`
        
        Conditions []Condition `json:"conditions"`
        
        Actions []Dialeraction `json:"actions"`
        *Alias
    }{

        


        


        


        


        
        Conditions: []Condition{{}},
        


        
        Actions: []Dialeraction{{}},
        

        Alias: (*Alias)(u),
    })
}

