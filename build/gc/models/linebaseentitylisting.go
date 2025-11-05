package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LinebaseentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LinebaseentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Linebaseentitylisting
type Linebaseentitylisting struct { 
    // Entities
    Entities []Linebase `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // TotalNumberOfEntities - The total organization-wide number of entities.
    TotalNumberOfEntities int `json:"totalNumberOfEntities"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Linebaseentitylisting) String() string {
     o.Entities = []Linebase{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Linebaseentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Linebaseentitylisting

    if LinebaseentitylistingMarshalled {
        return []byte("{}"), nil
    }
    LinebaseentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Linebase `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        TotalNumberOfEntities int `json:"totalNumberOfEntities"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Linebase{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

