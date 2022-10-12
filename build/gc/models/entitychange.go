package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EntitychangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EntitychangeDud struct { 
    


    


    


    


    

}

// Entitychange
type Entitychange struct { 
    // EntityId - Id of the entity that was changed
    EntityId string `json:"entityId"`


    // EntityName - Name of the entity that was changed
    EntityName string `json:"entityName"`


    // EntityType - Type of the entity that was changed
    EntityType string `json:"entityType"`


    // OldValues - Previous values for the entity.
    OldValues []string `json:"oldValues"`


    // NewValues - New values for the entity.
    NewValues []string `json:"newValues"`

}

// String returns a JSON representation of the model
func (o *Entitychange) String() string {
    
    
    
     o.OldValues = []string{""} 
     o.NewValues = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Entitychange) MarshalJSON() ([]byte, error) {
    type Alias Entitychange

    if EntitychangeMarshalled {
        return []byte("{}"), nil
    }
    EntitychangeMarshalled = true

    return json.Marshal(&struct {
        
        EntityId string `json:"entityId"`
        
        EntityName string `json:"entityName"`
        
        EntityType string `json:"entityType"`
        
        OldValues []string `json:"oldValues"`
        
        NewValues []string `json:"newValues"`
        *Alias
    }{

        


        


        


        
        OldValues: []string{""},
        


        
        NewValues: []string{""},
        

        Alias: (*Alias)(u),
    })
}

