package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextmessagelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextmessagelistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Textmessagelisting
type Textmessagelisting struct { 
    // Entities
    Entities []Messagedata `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // PageCount
    PageCount int `json:"pageCount"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`

}

// String returns a JSON representation of the model
func (o *Textmessagelisting) String() string {
     o.Entities = []Messagedata{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textmessagelisting) MarshalJSON() ([]byte, error) {
    type Alias Textmessagelisting

    if TextmessagelistingMarshalled {
        return []byte("{}"), nil
    }
    TextmessagelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Messagedata `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PageCount int `json:"pageCount"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        *Alias
    }{

        
        Entities: []Messagedata{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

