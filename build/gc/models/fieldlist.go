package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FieldlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FieldlistDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Fieldlist
type Fieldlist struct { 
    // CustomLabels
    CustomLabels bool `json:"customLabels"`


    // InstructionText
    InstructionText string `json:"instructionText"`


    // Key
    Key string `json:"key"`


    // LabelKeys
    LabelKeys []string `json:"labelKeys"`


    // Params
    Params map[string]interface{} `json:"params"`


    // Repeatable
    Repeatable bool `json:"repeatable"`


    // State
    State string `json:"state"`


    // VarType
    VarType string `json:"type"`


    // Required
    Required bool `json:"required"`


    // Gdpr
    Gdpr bool `json:"gdpr"`

}

// String returns a JSON representation of the model
func (o *Fieldlist) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.LabelKeys = []string{""} 
    
    
    
     o.Params = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fieldlist) MarshalJSON() ([]byte, error) {
    type Alias Fieldlist

    if FieldlistMarshalled {
        return []byte("{}"), nil
    }
    FieldlistMarshalled = true

    return json.Marshal(&struct { 
        CustomLabels bool `json:"customLabels"`
        
        InstructionText string `json:"instructionText"`
        
        Key string `json:"key"`
        
        LabelKeys []string `json:"labelKeys"`
        
        Params map[string]interface{} `json:"params"`
        
        Repeatable bool `json:"repeatable"`
        
        State string `json:"state"`
        
        VarType string `json:"type"`
        
        Required bool `json:"required"`
        
        Gdpr bool `json:"gdpr"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        LabelKeys: []string{""},
        

        

        
        Params: map[string]interface{}{"": Interface{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

