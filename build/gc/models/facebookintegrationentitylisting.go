package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookintegrationentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookintegrationentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Facebookintegrationentitylisting
type Facebookintegrationentitylisting struct { 
    // Entities
    Entities []Facebookintegration `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


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
func (o *Facebookintegrationentitylisting) String() string {
     o.Entities = []Facebookintegration{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookintegrationentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Facebookintegrationentitylisting

    if FacebookintegrationentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FacebookintegrationentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Facebookintegration `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Facebookintegration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

