package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypechangelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypechangelistingDud struct { 
    


    


    


    


    

}

// Worktypechangelisting
type Worktypechangelisting struct { 
    // Entities
    Entities []Workitemschangeworktypedelta `json:"entities"`


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
func (o *Worktypechangelisting) String() string {
     o.Entities = []Workitemschangeworktypedelta{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypechangelisting) MarshalJSON() ([]byte, error) {
    type Alias Worktypechangelisting

    if WorktypechangelistingMarshalled {
        return []byte("{}"), nil
    }
    WorktypechangelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemschangeworktypedelta `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemschangeworktypedelta{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

