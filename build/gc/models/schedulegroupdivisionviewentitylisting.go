package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegroupdivisionviewentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegroupdivisionviewentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Schedulegroupdivisionviewentitylisting
type Schedulegroupdivisionviewentitylisting struct { 
    // Entities
    Entities []Schedulegroupdivisionview `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Schedulegroupdivisionviewentitylisting) String() string {
     o.Entities = []Schedulegroupdivisionview{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegroupdivisionviewentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Schedulegroupdivisionviewentitylisting

    if SchedulegroupdivisionviewentitylistingMarshalled {
        return []byte("{}"), nil
    }
    SchedulegroupdivisionviewentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Schedulegroupdivisionview `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Schedulegroupdivisionview{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

