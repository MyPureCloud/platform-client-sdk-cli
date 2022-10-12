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
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Expandablewebdeployment `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

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
        
        Total int `json:"total"`
        
        Entities []Expandablewebdeployment `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Expandablewebdeployment{{}},
        


        

        Alias: (*Alias)(u),
    })
}

