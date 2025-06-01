package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookdataingestionruleversionresponseentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookdataingestionruleversionresponseentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Facebookdataingestionruleversionresponseentitylisting
type Facebookdataingestionruleversionresponseentitylisting struct { 
    // Entities
    Entities []Facebookdataingestionruleversionresponse `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // LastUri
    LastUri string `json:"lastUri"`


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
func (o *Facebookdataingestionruleversionresponseentitylisting) String() string {
     o.Entities = []Facebookdataingestionruleversionresponse{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookdataingestionruleversionresponseentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Facebookdataingestionruleversionresponseentitylisting

    if FacebookdataingestionruleversionresponseentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FacebookdataingestionruleversionresponseentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Facebookdataingestionruleversionresponse `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Facebookdataingestionruleversionresponse{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

