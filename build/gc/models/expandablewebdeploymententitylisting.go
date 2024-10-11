package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExpandablewebdeploymententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExpandablewebdeploymententitylistingDud struct { 
    


    


    


    


    

}

// Expandablewebdeploymententitylisting
type Expandablewebdeploymententitylisting struct { 
    // Entities
    Entities []Expandablewebdeployment `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // Total
    Total int `json:"total"`

}

// String returns a JSON representation of the model
func (o *Expandablewebdeploymententitylisting) String() string {
     o.Entities = []Expandablewebdeployment{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Expandablewebdeploymententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Expandablewebdeploymententitylisting

    if ExpandablewebdeploymententitylistingMarshalled {
        return []byte("{}"), nil
    }
    ExpandablewebdeploymententitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Expandablewebdeployment `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Total int `json:"total"`
        *Alias
    }{

        
        Entities: []Expandablewebdeployment{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

