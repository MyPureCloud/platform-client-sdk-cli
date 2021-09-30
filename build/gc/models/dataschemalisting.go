package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataschemalistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataschemalistingDud struct { 
    


    


    

}

// Dataschemalisting
type Dataschemalisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Dataschema `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Dataschemalisting) String() string {
    
    
    
    
    
    
     o.Entities = []Dataschema{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataschemalisting) MarshalJSON() ([]byte, error) {
    type Alias Dataschemalisting

    if DataschemalistingMarshalled {
        return []byte("{}"), nil
    }
    DataschemalistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Dataschema `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Dataschema{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

