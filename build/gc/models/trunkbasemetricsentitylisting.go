package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkbasemetricsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkbasemetricsentitylistingDud struct { 
    


    


    

}

// Trunkbasemetricsentitylisting
type Trunkbasemetricsentitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Trunkbasemetrics `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Trunkbasemetricsentitylisting) String() string {
    
     o.Entities = []Trunkbasemetrics{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkbasemetricsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Trunkbasemetricsentitylisting

    if TrunkbasemetricsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TrunkbasemetricsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Trunkbasemetrics `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Trunkbasemetrics{{}},
        


        

        Alias: (*Alias)(u),
    })
}

