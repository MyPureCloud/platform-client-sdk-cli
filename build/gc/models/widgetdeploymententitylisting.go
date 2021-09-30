package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WidgetdeploymententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WidgetdeploymententitylistingDud struct { 
    


    


    

}

// Widgetdeploymententitylisting
type Widgetdeploymententitylisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Widgetdeployment `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Widgetdeploymententitylisting) String() string {
    
    
    
    
    
    
     o.Entities = []Widgetdeployment{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Widgetdeploymententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Widgetdeploymententitylisting

    if WidgetdeploymententitylistingMarshalled {
        return []byte("{}"), nil
    }
    WidgetdeploymententitylistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Widgetdeployment `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Widgetdeployment{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

