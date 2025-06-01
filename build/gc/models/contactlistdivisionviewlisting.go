package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistdivisionviewlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistdivisionviewlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Contactlistdivisionviewlisting
type Contactlistdivisionviewlisting struct { 
    // Entities
    Entities []Contactlistdivisionview `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // LastUri
    LastUri string `json:"lastUri"`


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
func (o *Contactlistdivisionviewlisting) String() string {
     o.Entities = []Contactlistdivisionview{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistdivisionviewlisting) MarshalJSON() ([]byte, error) {
    type Alias Contactlistdivisionviewlisting

    if ContactlistdivisionviewlistingMarshalled {
        return []byte("{}"), nil
    }
    ContactlistdivisionviewlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactlistdivisionview `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Contactlistdivisionview{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

