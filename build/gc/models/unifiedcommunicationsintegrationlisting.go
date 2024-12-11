package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnifiedcommunicationsintegrationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnifiedcommunicationsintegrationlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Unifiedcommunicationsintegrationlisting
type Unifiedcommunicationsintegrationlisting struct { 
    // Entities
    Entities []Unifiedcommunicationsintegration `json:"entities"`


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
func (o *Unifiedcommunicationsintegrationlisting) String() string {
     o.Entities = []Unifiedcommunicationsintegration{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unifiedcommunicationsintegrationlisting) MarshalJSON() ([]byte, error) {
    type Alias Unifiedcommunicationsintegrationlisting

    if UnifiedcommunicationsintegrationlistingMarshalled {
        return []byte("{}"), nil
    }
    UnifiedcommunicationsintegrationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Unifiedcommunicationsintegration `json:"entities"`
        
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

        
        Entities: []Unifiedcommunicationsintegration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

