package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CommoncampaigndivisionviewentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CommoncampaigndivisionviewentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Commoncampaigndivisionviewentitylisting
type Commoncampaigndivisionviewentitylisting struct { 
    // Entities
    Entities []Commoncampaigndivisionview `json:"entities"`


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
func (o *Commoncampaigndivisionviewentitylisting) String() string {
    
    
     o.Entities = []Commoncampaigndivisionview{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Commoncampaigndivisionviewentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Commoncampaigndivisionviewentitylisting

    if CommoncampaigndivisionviewentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CommoncampaigndivisionviewentitylistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Commoncampaigndivisionview `json:"entities"`
        
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
        

        
        Entities: []Commoncampaigndivisionview{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

