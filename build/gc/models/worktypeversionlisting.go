package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypeversionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypeversionlistingDud struct { 
    


    


    


    


    

}

// Worktypeversionlisting
type Worktypeversionlisting struct { 
    // Entities
    Entities []Worktypeversion `json:"entities"`


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
func (o *Worktypeversionlisting) String() string {
     o.Entities = []Worktypeversion{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypeversionlisting) MarshalJSON() ([]byte, error) {
    type Alias Worktypeversionlisting

    if WorktypeversionlistingMarshalled {
        return []byte("{}"), nil
    }
    WorktypeversionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Worktypeversion `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Worktypeversion{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

