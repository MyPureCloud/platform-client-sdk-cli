package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmaplistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmaplistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Actionmaplisting
type Actionmaplisting struct { 
    // Entities
    Entities []Actionmap `json:"entities"`


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
func (o *Actionmaplisting) String() string {
     o.Entities = []Actionmap{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmaplisting) MarshalJSON() ([]byte, error) {
    type Alias Actionmaplisting

    if ActionmaplistingMarshalled {
        return []byte("{}"), nil
    }
    ActionmaplistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Actionmap `json:"entities"`
        
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

        
        Entities: []Actionmap{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

