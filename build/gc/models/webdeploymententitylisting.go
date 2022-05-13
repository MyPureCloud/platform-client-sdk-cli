package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymententitylistingDud struct { 
    


    


    

}

// Webdeploymententitylisting
type Webdeploymententitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Webdeployment `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Webdeploymententitylisting) String() string {
    
     o.Entities = []Webdeployment{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymententitylisting

    if WebdeploymententitylistingMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymententitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Webdeployment `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Webdeployment{{}},
        


        

        Alias: (*Alias)(u),
    })
}

