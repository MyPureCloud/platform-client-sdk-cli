package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnoredminedtopiclistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnoredminedtopiclistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Ignoredminedtopiclisting
type Ignoredminedtopiclisting struct { 
    // Entities
    Entities []Ignoredminedentity `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Ignoredminedtopiclisting) String() string {
     o.Entities = []Ignoredminedentity{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignoredminedtopiclisting) MarshalJSON() ([]byte, error) {
    type Alias Ignoredminedtopiclisting

    if IgnoredminedtopiclistingMarshalled {
        return []byte("{}"), nil
    }
    IgnoredminedtopiclistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Ignoredminedentity `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Ignoredminedentity{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

