package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResourceconditionnodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResourceconditionnodeDud struct { 
    


    


    


    


    

}

// Resourceconditionnode
type Resourceconditionnode struct { 
    // VariableName
    VariableName string `json:"variableName"`


    // Conjunction
    Conjunction string `json:"conjunction"`


    // Operator
    Operator string `json:"operator"`


    // Operands
    Operands []Resourceconditionvalue `json:"operands"`


    // Terms
    Terms []Resourceconditionnode `json:"terms"`

}

// String returns a JSON representation of the model
func (o *Resourceconditionnode) String() string {
    
    
    
     o.Operands = []Resourceconditionvalue{{}} 
     o.Terms = []Resourceconditionnode{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resourceconditionnode) MarshalJSON() ([]byte, error) {
    type Alias Resourceconditionnode

    if ResourceconditionnodeMarshalled {
        return []byte("{}"), nil
    }
    ResourceconditionnodeMarshalled = true

    return json.Marshal(&struct {
        
        VariableName string `json:"variableName"`
        
        Conjunction string `json:"conjunction"`
        
        Operator string `json:"operator"`
        
        Operands []Resourceconditionvalue `json:"operands"`
        
        Terms []Resourceconditionnode `json:"terms"`
        *Alias
    }{

        


        


        


        
        Operands: []Resourceconditionvalue{{}},
        


        
        Terms: []Resourceconditionnode{{}},
        

        Alias: (*Alias)(u),
    })
}

