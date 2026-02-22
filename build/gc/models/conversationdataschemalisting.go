package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationdataschemalistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationdataschemalistingDud struct { 
    


    


    

}

// Conversationdataschemalisting
type Conversationdataschemalisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Conversationdataschema `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Conversationdataschemalisting) String() string {
    
     o.Entities = []Conversationdataschema{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationdataschemalisting) MarshalJSON() ([]byte, error) {
    type Alias Conversationdataschemalisting

    if ConversationdataschemalistingMarshalled {
        return []byte("{}"), nil
    }
    ConversationdataschemalistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Conversationdataschema `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Conversationdataschema{{}},
        


        

        Alias: (*Alias)(u),
    })
}

