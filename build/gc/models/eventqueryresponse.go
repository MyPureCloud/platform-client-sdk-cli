package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventqueryresponseDud struct { 
    


    


    


    

}

// Eventqueryresponse
type Eventqueryresponse struct { 
    // Entities
    Entities []Operationalevent `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Eventqueryresponse) String() string {
     o.Entities = []Operationalevent{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Eventqueryresponse

    if EventqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    EventqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Operationalevent `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Operationalevent{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}
