package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgegroupentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgegroupentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Edgegroupentitylisting
type Edgegroupentitylisting struct { 
    // Entities
    Entities []Edgegroup `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Edgegroupentitylisting) String() string {
     o.Entities = []Edgegroup{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgegroupentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Edgegroupentitylisting

    if EdgegroupentitylistingMarshalled {
        return []byte("{}"), nil
    }
    EdgegroupentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Edgegroup `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Edgegroup{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

