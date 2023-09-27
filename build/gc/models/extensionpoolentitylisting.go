package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExtensionpoolentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExtensionpoolentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Extensionpoolentitylisting
type Extensionpoolentitylisting struct { 
    // Entities
    Entities []Extensionpool `json:"entities"`


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


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Extensionpoolentitylisting) String() string {
     o.Entities = []Extensionpool{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Extensionpoolentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Extensionpoolentitylisting

    if ExtensionpoolentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ExtensionpoolentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Extensionpool `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Extensionpool{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

