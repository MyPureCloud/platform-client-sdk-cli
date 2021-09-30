package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvententityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvententityDud struct { 
    


    

}

// Evententity
type Evententity struct { 
    // EntityType - Type of entity the event pertains to. e.g. integration
    EntityType string `json:"entityType"`


    // Id - ID of the entity the event pertains to.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Evententity) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evententity) MarshalJSON() ([]byte, error) {
    type Alias Evententity

    if EvententityMarshalled {
        return []byte("{}"), nil
    }
    EvententityMarshalled = true

    return json.Marshal(&struct { 
        EntityType string `json:"entityType"`
        
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

