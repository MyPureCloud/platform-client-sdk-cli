package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramsmappingsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramsmappingsentitylistingDud struct { 
    


    


    


    


    

}

// Programsmappingsentitylisting
type Programsmappingsentitylisting struct { 
    // Entities
    Entities []Programmappings `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Programsmappingsentitylisting) String() string {
     o.Entities = []Programmappings{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programsmappingsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Programsmappingsentitylisting

    if ProgramsmappingsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ProgramsmappingsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Programmappings `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        
        Entities: []Programmappings{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

