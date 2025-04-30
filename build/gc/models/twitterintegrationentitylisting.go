package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterintegrationentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterintegrationentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Twitterintegrationentitylisting
type Twitterintegrationentitylisting struct { 
    // Entities
    Entities []Twitterintegration `json:"entities"`


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
func (o *Twitterintegrationentitylisting) String() string {
     o.Entities = []Twitterintegration{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterintegrationentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Twitterintegrationentitylisting

    if TwitterintegrationentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TwitterintegrationentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Twitterintegration `json:"entities"`
        
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

        
        Entities: []Twitterintegration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

