package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotconnectorbotsummaryentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotconnectorbotsummaryentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Botconnectorbotsummaryentitylisting
type Botconnectorbotsummaryentitylisting struct { 
    // Entities
    Entities []Botsummary `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Botconnectorbotsummaryentitylisting) String() string {
     o.Entities = []Botsummary{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botconnectorbotsummaryentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Botconnectorbotsummaryentitylisting

    if BotconnectorbotsummaryentitylistingMarshalled {
        return []byte("{}"), nil
    }
    BotconnectorbotsummaryentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Botsummary `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Botsummary{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

