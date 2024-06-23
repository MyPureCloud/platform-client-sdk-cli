package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkbaseentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkbaseentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Trunkbaseentitylisting
type Trunkbaseentitylisting struct { 
    // Entities
    Entities []Trunkbase `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


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
func (o *Trunkbaseentitylisting) String() string {
     o.Entities = []Trunkbase{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkbaseentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Trunkbaseentitylisting

    if TrunkbaseentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TrunkbaseentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Trunkbase `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Trunkbase{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

