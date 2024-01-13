package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PagelessentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PagelessentitylistingDud struct { 
    


    


    

}

// Pagelessentitylisting
type Pagelessentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Addressableentity `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Pagelessentitylisting) String() string {
    
     o.Entities = []Addressableentity{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pagelessentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Pagelessentitylisting

    if PagelessentitylistingMarshalled {
        return []byte("{}"), nil
    }
    PagelessentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Addressableentity `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Addressableentity{{}},
        


        

        Alias: (*Alias)(u),
    })
}

