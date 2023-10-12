package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypequeryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypequeryentitylistingDud struct { 
    


    


    


    


    


    

}

// Worktypequeryentitylisting
type Worktypequeryentitylisting struct { 
    // Entities
    Entities []Worktype `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`


    // Count - Count of items returned by the query. Refer to the \"select\" request parameter for more information.
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Worktypequeryentitylisting) String() string {
     o.Entities = []Worktype{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypequeryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Worktypequeryentitylisting

    if WorktypequeryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    WorktypequeryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Worktype `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        
        Count int `json:"count"`
        *Alias
    }{

        
        Entities: []Worktype{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

