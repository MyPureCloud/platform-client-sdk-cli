package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowoutcomedivisionviewentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowoutcomedivisionviewentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Flowoutcomedivisionviewentitylisting
type Flowoutcomedivisionviewentitylisting struct { 
    // Entities
    Entities []Flowoutcomedivisionview `json:"entities"`


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


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Flowoutcomedivisionviewentitylisting) String() string {
     o.Entities = []Flowoutcomedivisionview{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowoutcomedivisionviewentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Flowoutcomedivisionviewentitylisting

    if FlowoutcomedivisionviewentitylistingMarshalled {
        return []byte("{}"), nil
    }
    FlowoutcomedivisionviewentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Flowoutcomedivisionview `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Flowoutcomedivisionview{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

