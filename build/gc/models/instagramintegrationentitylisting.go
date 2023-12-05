package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramintegrationentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramintegrationentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Instagramintegrationentitylisting
type Instagramintegrationentitylisting struct { 
    // Entities
    Entities []Instagramintegration `json:"entities"`


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


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Instagramintegrationentitylisting) String() string {
     o.Entities = []Instagramintegration{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramintegrationentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Instagramintegrationentitylisting

    if InstagramintegrationentitylistingMarshalled {
        return []byte("{}"), nil
    }
    InstagramintegrationentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Instagramintegration `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Instagramintegration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

