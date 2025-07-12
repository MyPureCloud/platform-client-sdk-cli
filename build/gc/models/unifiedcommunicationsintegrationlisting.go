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
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Unifiedcommunicationsintegration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

