package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamentitylistingDud struct { 
    


    


    


    

}

// Teamentitylisting
type Teamentitylisting struct { 
    // Entities
    Entities []Team `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Teamentitylisting) String() string {
     o.Entities = []Team{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Teamentitylisting

    if TeamentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TeamentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Team `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Team{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

