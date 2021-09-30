package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainresourceconditionnodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainresourceconditionnodeDud struct { 
    


    


    


    


    

}

// Domainresourceconditionnode
type Domainresourceconditionnode struct { 
    // VariableName
    VariableName string `json:"variableName"`


    // Operator
    Operator string `json:"operator"`


    // Operands
    Operands []Domainresourceconditionvalue `json:"operands"`


    // Conjunction
    Conjunction string `json:"conjunction"`


    // Terms
    Terms []Domainresourceconditionnode `json:"terms"`

}

// String returns a JSON representation of the model
func (o *Domainresourceconditionnode) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Operands = []Domainresourceconditionvalue{{}} 
    
    
    
    
    
    
    
     o.Terms = []Domainresourceconditionnode{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainresourceconditionnode) MarshalJSON() ([]byte, error) {
    type Alias Domainresourceconditionnode

    if DomainresourceconditionnodeMarshalled {
        return []byte("{}"), nil
    }
    DomainresourceconditionnodeMarshalled = true

    return json.Marshal(&struct { 
        VariableName string `json:"variableName"`
        
        Operator string `json:"operator"`
        
        Operands []Domainresourceconditionvalue `json:"operands"`
        
        Conjunction string `json:"conjunction"`
        
        Terms []Domainresourceconditionnode `json:"terms"`
        
        *Alias
    }{
        

        

        

        

        

        
        Operands: []Domainresourceconditionvalue{{}},
        

        

        

        

        
        Terms: []Domainresourceconditionnode{{}},
        

        
        Alias: (*Alias)(u),
    })
}

