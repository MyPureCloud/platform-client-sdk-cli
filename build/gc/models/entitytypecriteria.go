package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EntitytypecriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EntitytypecriteriaDud struct { 
    


    


    


    


    

}

// Entitytypecriteria
type Entitytypecriteria struct { 
    // Key - The criteria key.
    Key string `json:"key"`


    // Values - The criteria values.
    Values []string `json:"values"`


    // ShouldIgnoreCase - Should criteria be case insensitive.
    ShouldIgnoreCase bool `json:"shouldIgnoreCase"`


    // Operator - The comparison operator.
    Operator string `json:"operator"`


    // EntityType - The entity to match the pattern against.
    EntityType string `json:"entityType"`

}

// String returns a JSON representation of the model
func (o *Entitytypecriteria) String() string {
    
    
    
    
    
    
     o.Values = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Entitytypecriteria) MarshalJSON() ([]byte, error) {
    type Alias Entitytypecriteria

    if EntitytypecriteriaMarshalled {
        return []byte("{}"), nil
    }
    EntitytypecriteriaMarshalled = true

    return json.Marshal(&struct { 
        Key string `json:"key"`
        
        Values []string `json:"values"`
        
        ShouldIgnoreCase bool `json:"shouldIgnoreCase"`
        
        Operator string `json:"operator"`
        
        EntityType string `json:"entityType"`
        
        *Alias
    }{
        

        

        

        
        Values: []string{""},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

