package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangerulelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangerulelistingDud struct { 
    


    


    


    


    

}

// Workitemonattributechangerulelisting
type Workitemonattributechangerulelisting struct { 
    // Entities
    Entities []Workitemonattributechangerule `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangerulelisting) String() string {
     o.Entities = []Workitemonattributechangerule{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangerulelisting) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangerulelisting

    if WorkitemonattributechangerulelistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangerulelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitemonattributechangerule `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        *Alias
    }{

        
        Entities: []Workitemonattributechangerule{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

