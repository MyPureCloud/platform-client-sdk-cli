package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitempostqueryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitempostqueryentitylistingDud struct { 
    


    


    


    


    


    

}

// Workitempostqueryentitylisting
type Workitempostqueryentitylisting struct { 
    // Entities
    Entities []Workitem `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // After
    After string `json:"after"`


    // Count
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Workitempostqueryentitylisting) String() string {
     o.Entities = []Workitem{{}} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitempostqueryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Workitempostqueryentitylisting

    if WorkitempostqueryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    WorkitempostqueryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Workitem `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        After string `json:"after"`
        
        Count int `json:"count"`
        *Alias
    }{

        
        Entities: []Workitem{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

