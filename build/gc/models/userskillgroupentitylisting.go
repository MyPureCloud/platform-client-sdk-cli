package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserskillgroupentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserskillgroupentitylistingDud struct { 
    


    


    


    

}

// Userskillgroupentitylisting
type Userskillgroupentitylisting struct { 
    // Entities
    Entities []Skillgroup `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Userskillgroupentitylisting) String() string {
     o.Entities = []Skillgroup{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userskillgroupentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Userskillgroupentitylisting

    if UserskillgroupentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UserskillgroupentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Skillgroup `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Skillgroup{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

