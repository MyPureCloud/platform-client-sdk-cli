package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GetcontestsessentialslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GetcontestsessentialslistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Getcontestsessentialslisting
type Getcontestsessentialslisting struct { 
    // Entities
    Entities []Contestsessentials `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


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
func (o *Getcontestsessentialslisting) String() string {
     o.Entities = []Contestsessentials{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Getcontestsessentialslisting) MarshalJSON() ([]byte, error) {
    type Alias Getcontestsessentialslisting

    if GetcontestsessentialslistingMarshalled {
        return []byte("{}"), nil
    }
    GetcontestsessentialslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Contestsessentials `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Contestsessentials{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

