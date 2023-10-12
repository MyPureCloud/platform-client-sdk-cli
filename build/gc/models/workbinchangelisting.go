package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkbinchangelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkbinchangelistingDud struct { 
    


    


    


    


    

}

// Workbinchangelisting
type Workbinchangelisting struct { 
    // Entities
    Entities []Workitemschangeworkbindelta `json:"entities"`


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
func (o *Workbinchangelisting) String() string {
     o.Entities = []Workitemschangeworkbindelta{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workbinchangelisting) MarshalJSON() ([]byte, error) {
    type Alias Workbinchangelisting

    if WorkbinchangelistingMarshalled {
        return []byte("{}"), nil
    }
    WorkbinchangelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemschangeworkbindelta `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemschangeworkbindelta{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

