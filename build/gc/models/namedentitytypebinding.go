package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentitytypebindingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentitytypebindingDud struct { 
    


    

}

// Namedentitytypebinding
type Namedentitytypebinding struct { 
    // EntityType - The named entity type of the binding. It can be a built-in one such as builtin:number or a custom entity type such as BeverageType.
    EntityType string `json:"entityType"`


    // EntityName - The name that this named entity type is bound to.
    EntityName string `json:"entityName"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypebinding) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentitytypebinding) MarshalJSON() ([]byte, error) {
    type Alias Namedentitytypebinding

    if NamedentitytypebindingMarshalled {
        return []byte("{}"), nil
    }
    NamedentitytypebindingMarshalled = true

    return json.Marshal(&struct { 
        EntityType string `json:"entityType"`
        
        EntityName string `json:"entityName"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

