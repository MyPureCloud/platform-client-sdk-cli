package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotruleDud struct { 
    


    

}

// Copilotrule
type Copilotrule struct { 
    // Conditions - List of conditions to execute actions, must have at least 1 element and maximum 100 elements. Operator in case of multiple conditions: 'OR'.
    Conditions []Copilotcondition `json:"conditions"`


    // Actions - List of actions to execute, must have at least 1 element and maximum 100 elements.
    Actions []Copilotaction `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Copilotrule) String() string {
     o.Conditions = []Copilotcondition{{}} 
     o.Actions = []Copilotaction{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilotrule) MarshalJSON() ([]byte, error) {
    type Alias Copilotrule

    if CopilotruleMarshalled {
        return []byte("{}"), nil
    }
    CopilotruleMarshalled = true

    return json.Marshal(&struct {
        
        Conditions []Copilotcondition `json:"conditions"`
        
        Actions []Copilotaction `json:"actions"`
        *Alias
    }{

        
        Conditions: []Copilotcondition{{}},
        


        
        Actions: []Copilotaction{{}},
        

        Alias: (*Alias)(u),
    })
}

