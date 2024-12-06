package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IvrdivisionviewentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IvrdivisionviewentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Ivrdivisionviewentitylisting
type Ivrdivisionviewentitylisting struct { 
    // Entities
    Entities []Ivrdivisionview `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // TotalNumberOfEntities - The total organization-wide number of entities.
    TotalNumberOfEntities int `json:"totalNumberOfEntities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Ivrdivisionviewentitylisting) String() string {
     o.Entities = []Ivrdivisionview{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ivrdivisionviewentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Ivrdivisionviewentitylisting

    if IvrdivisionviewentitylistingMarshalled {
        return []byte("{}"), nil
    }
    IvrdivisionviewentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Ivrdivisionview `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        TotalNumberOfEntities int `json:"totalNumberOfEntities"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Ivrdivisionview{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

