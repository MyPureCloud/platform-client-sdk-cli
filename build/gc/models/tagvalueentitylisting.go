package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TagvalueentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TagvalueentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Tagvalueentitylisting
type Tagvalueentitylisting struct { 
    // Entities
    Entities []Tagvalue `json:"entities"`


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
func (o *Tagvalueentitylisting) String() string {
    
    
     o.Entities = []Tagvalue{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Tagvalueentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Tagvalueentitylisting

    if TagvalueentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TagvalueentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Tagvalue `json:"entities"`
        
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
        

        
        Entities: []Tagvalue{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

