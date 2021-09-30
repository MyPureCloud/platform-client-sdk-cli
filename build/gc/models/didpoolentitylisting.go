package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DidpoolentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DidpoolentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Didpoolentitylisting
type Didpoolentitylisting struct { 
    // Entities
    Entities []Didpool `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Didpoolentitylisting) String() string {
    
    
     o.Entities = []Didpool{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Didpoolentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Didpoolentitylisting

    if DidpoolentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DidpoolentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Didpool `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Didpool{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

