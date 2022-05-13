package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FieldconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FieldconfigDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Fieldconfig
type Fieldconfig struct { 
    


    // Name
    Name string `json:"name"`


    // EntityType
    EntityType string `json:"entityType"`


    // State
    State string `json:"state"`


    // Sections
    Sections []Section `json:"sections"`


    // Version
    Version string `json:"version"`


    // SchemaVersion
    SchemaVersion string `json:"schemaVersion"`


    

}

// String returns a JSON representation of the model
func (o *Fieldconfig) String() string {
    
    
    
     o.Sections = []Section{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fieldconfig) MarshalJSON() ([]byte, error) {
    type Alias Fieldconfig

    if FieldconfigMarshalled {
        return []byte("{}"), nil
    }
    FieldconfigMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        EntityType string `json:"entityType"`
        
        State string `json:"state"`
        
        Sections []Section `json:"sections"`
        
        Version string `json:"version"`
        
        SchemaVersion string `json:"schemaVersion"`
        *Alias
    }{

        


        


        


        


        
        Sections: []Section{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

