package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlisttemplateentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlisttemplateentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Contactlisttemplateentitylisting
type Contactlisttemplateentitylisting struct { 
    // Entities
    Entities []Contactlisttemplate `json:"entities"`


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


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Contactlisttemplateentitylisting) String() string {
     o.Entities = []Contactlisttemplate{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlisttemplateentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Contactlisttemplateentitylisting

    if ContactlisttemplateentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ContactlisttemplateentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contactlisttemplate `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Contactlisttemplate{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

