package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupentitylistingDud struct { 
    


    


    


    

}

// Skillgroupentitylisting
type Skillgroupentitylisting struct { 
    // Entities
    Entities []Skillgroupdefinition `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Skillgroupentitylisting) String() string {
     o.Entities = []Skillgroupdefinition{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroupentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Skillgroupentitylisting

    if SkillgroupentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Skillgroupdefinition `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Skillgroupdefinition{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

