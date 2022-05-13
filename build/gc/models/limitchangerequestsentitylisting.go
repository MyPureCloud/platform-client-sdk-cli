package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LimitchangerequestsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LimitchangerequestsentitylistingDud struct { 
    


    


    


    

}

// Limitchangerequestsentitylisting
type Limitchangerequestsentitylisting struct { 
    // Entities
    Entities []Limitchangerequestdetails `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Limitchangerequestsentitylisting) String() string {
     o.Entities = []Limitchangerequestdetails{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Limitchangerequestsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Limitchangerequestsentitylisting

    if LimitchangerequestsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    LimitchangerequestsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Limitchangerequestdetails `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Limitchangerequestdetails{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

