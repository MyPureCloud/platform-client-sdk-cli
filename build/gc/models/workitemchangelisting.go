package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemchangelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemchangelistingDud struct { 
    


    


    


    


    

}

// Workitemchangelisting
type Workitemchangelisting struct { 
    // Entities
    Entities []Workitemschangeworkitemdelta `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Workitemchangelisting) String() string {
     o.Entities = []Workitemschangeworkitemdelta{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemchangelisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemchangelisting

    if WorkitemchangelistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemchangelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemschangeworkitemdelta `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemschangeworkitemdelta{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

