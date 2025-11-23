package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowcxagentsummaryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowcxagentsummaryentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Dialogflowcxagentsummaryentitylisting
type Dialogflowcxagentsummaryentitylisting struct { 
    // Entities
    Entities []Dialogflowcxagentsummary `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Dialogflowcxagentsummaryentitylisting) String() string {
     o.Entities = []Dialogflowcxagentsummary{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowcxagentsummaryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowcxagentsummaryentitylisting

    if DialogflowcxagentsummaryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    DialogflowcxagentsummaryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dialogflowcxagentsummary `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Dialogflowcxagentsummary{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

