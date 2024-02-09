package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchentitytypecriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchentitytypecriteriaDud struct { 
    


    


    


    


    

}

// Patchentitytypecriteria
type Patchentitytypecriteria struct { 
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
func (o *Patchentitytypecriteria) String() string {
    
     o.Values = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchentitytypecriteria) MarshalJSON() ([]byte, error) {
    type Alias Patchentitytypecriteria

    if PatchentitytypecriteriaMarshalled {
        return []byte("{}"), nil
    }
    PatchentitytypecriteriaMarshalled = true

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

