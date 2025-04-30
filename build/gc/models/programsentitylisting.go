package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramsentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramsentitylistingDud struct { 
    


    


    


    


    

}

// Programsentitylisting
type Programsentitylisting struct { 
    // Entities
    Entities []Listedprogram `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Programsentitylisting) String() string {
     o.Entities = []Listedprogram{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programsentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Programsentitylisting

    if ProgramsentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ProgramsentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Listedprogram `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Listedprogram{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

