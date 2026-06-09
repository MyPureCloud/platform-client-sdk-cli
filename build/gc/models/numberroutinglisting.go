package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumberroutinglistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumberroutinglistingDud struct { 
    


    


    


    

}

// Numberroutinglisting
type Numberroutinglisting struct { 
    // Entities
    Entities []Numberrouting `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Numberroutinglisting) String() string {
     o.Entities = []Numberrouting{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Numberroutinglisting) MarshalJSON() ([]byte, error) {
    type Alias Numberroutinglisting

    if NumberroutinglistingMarshalled {
        return []byte("{}"), nil
    }
    NumberroutinglistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Numberrouting `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Numberrouting{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

