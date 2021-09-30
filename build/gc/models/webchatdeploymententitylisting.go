package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatdeploymententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatdeploymententitylistingDud struct { 
    


    


    

}

// Webchatdeploymententitylisting
type Webchatdeploymententitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Webchatdeployment `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Webchatdeploymententitylisting) String() string {
    
    
    
    
    
    
     o.Entities = []Webchatdeployment{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatdeploymententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Webchatdeploymententitylisting

    if WebchatdeploymententitylistingMarshalled {
        return []byte("{}"), nil
    }
    WebchatdeploymententitylistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Webchatdeployment `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Webchatdeployment{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

