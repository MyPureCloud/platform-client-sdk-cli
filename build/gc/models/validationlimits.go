package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidationlimitsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidationlimitsDud struct { 
    


    


    


    


    


    

}

// Validationlimits
type Validationlimits struct { 
    // MinLength
    MinLength Minlength `json:"minLength"`


    // MaxLength
    MaxLength Maxlength `json:"maxLength"`


    // MinItems
    MinItems Minlength `json:"minItems"`


    // MaxItems
    MaxItems Maxlength `json:"maxItems"`


    // Minimum
    Minimum Minlength `json:"minimum"`


    // Maximum
    Maximum Maxlength `json:"maximum"`

}

// String returns a JSON representation of the model
func (o *Validationlimits) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validationlimits) MarshalJSON() ([]byte, error) {
    type Alias Validationlimits

    if ValidationlimitsMarshalled {
        return []byte("{}"), nil
    }
    ValidationlimitsMarshalled = true

    return json.Marshal(&struct { 
        MinLength Minlength `json:"minLength"`
        
        MaxLength Maxlength `json:"maxLength"`
        
        MinItems Minlength `json:"minItems"`
        
        MaxItems Maxlength `json:"maxItems"`
        
        Minimum Minlength `json:"minimum"`
        
        Maximum Maxlength `json:"maximum"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

