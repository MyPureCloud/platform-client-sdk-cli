package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseentitylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseentitylistDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Responseentitylist - Query result list
type Responseentitylist struct { 
    // Entities
    Entities []Response `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount - Total number of pages
    PageCount int `json:"pageCount"`


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

}

// String returns a JSON representation of the model
func (o *Responseentitylist) String() string {
     o.Entities = []Response{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseentitylist) MarshalJSON() ([]byte, error) {
    type Alias Responseentitylist

    if ResponseentitylistMarshalled {
        return []byte("{}"), nil
    }
    ResponseentitylistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Response `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        
        Entities: []Response{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

