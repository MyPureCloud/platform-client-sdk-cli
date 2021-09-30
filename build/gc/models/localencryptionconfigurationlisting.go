package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocalencryptionconfigurationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocalencryptionconfigurationlistingDud struct { 
    


    


    

}

// Localencryptionconfigurationlisting
type Localencryptionconfigurationlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Localencryptionconfiguration `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Localencryptionconfigurationlisting) String() string {
    
    
    
    
    
    
     o.Entities = []Localencryptionconfiguration{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localencryptionconfigurationlisting) MarshalJSON() ([]byte, error) {
    type Alias Localencryptionconfigurationlisting

    if LocalencryptionconfigurationlistingMarshalled {
        return []byte("{}"), nil
    }
    LocalencryptionconfigurationlistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Localencryptionconfiguration `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Localencryptionconfiguration{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

