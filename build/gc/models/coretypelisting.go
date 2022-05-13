package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoretypelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoretypelistingDud struct { 
    


    


    

}

// Coretypelisting
type Coretypelisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Coretype `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Coretypelisting) String() string {
    
     o.Entities = []Coretype{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coretypelisting) MarshalJSON() ([]byte, error) {
    type Alias Coretypelisting

    if CoretypelistingMarshalled {
        return []byte("{}"), nil
    }
    CoretypelistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Coretype `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Coretype{{}},
        


        

        Alias: (*Alias)(u),
    })
}

