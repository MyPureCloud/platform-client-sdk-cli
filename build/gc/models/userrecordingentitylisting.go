package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserrecordingentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserrecordingentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Userrecordingentitylisting
type Userrecordingentitylisting struct { 
    // Entities
    Entities []Userrecording `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


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
func (o *Userrecordingentitylisting) String() string {
     o.Entities = []Userrecording{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userrecordingentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Userrecordingentitylisting

    if UserrecordingentitylistingMarshalled {
        return []byte("{}"), nil
    }
    UserrecordingentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Userrecording `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Userrecording{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

