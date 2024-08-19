package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IvrentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IvrentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Ivrentitylisting
type Ivrentitylisting struct { 
    // Entities
    Entities []Ivr `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // TotalNumberOfEntities - The total organization-wide number of entities.
    TotalNumberOfEntities int `json:"totalNumberOfEntities"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Ivrentitylisting) String() string {
     o.Entities = []Ivr{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ivrentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Ivrentitylisting

    if IvrentitylistingMarshalled {
        return []byte("{}"), nil
    }
    IvrentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Ivr `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        TotalNumberOfEntities int `json:"totalNumberOfEntities"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Ivr{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

