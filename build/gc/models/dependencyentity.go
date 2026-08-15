package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DependencyentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DependencyentityDud struct { 
    


    

}

// Dependencyentity - A dependency entity with its type and ID.
type Dependencyentity struct { 
    // EntityId - The ID of the entity.
    EntityId string `json:"entityId"`


    // EntityType - The type of the entity
    EntityType string `json:"entityType"`

}

// String returns a JSON representation of the model
func (o *Dependencyentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dependencyentity) MarshalJSON() ([]byte, error) {
    type Alias Dependencyentity

    if DependencyentityMarshalled {
        return []byte("{}"), nil
    }
    DependencyentityMarshalled = true

    return json.Marshal(&struct {
        
        EntityId string `json:"entityId"`
        
        EntityType string `json:"entityType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

