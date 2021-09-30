package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnsresponseDud struct { 
    


    


    


    

}

// Reportingturnsresponse
type Reportingturnsresponse struct { 
    // Entities
    Entities []Reportingturn `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Reportingturnsresponse) String() string {
    
    
     o.Entities = []Reportingturn{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnsresponse) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnsresponse

    if ReportingturnsresponseMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnsresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Reportingturn `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        *Alias
    }{
        

        
        Entities: []Reportingturn{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

