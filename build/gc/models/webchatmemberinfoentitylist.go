package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatmemberinfoentitylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatmemberinfoentitylistDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Webchatmemberinfoentitylist
type Webchatmemberinfoentitylist struct { 
    // Entities
    Entities []Webchatmemberinfo `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Webchatmemberinfoentitylist) String() string {
     o.Entities = []Webchatmemberinfo{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatmemberinfoentitylist) MarshalJSON() ([]byte, error) {
    type Alias Webchatmemberinfoentitylist

    if WebchatmemberinfoentitylistMarshalled {
        return []byte("{}"), nil
    }
    WebchatmemberinfoentitylistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Webchatmemberinfo `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Webchatmemberinfo{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

