package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CobrowseconversationentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CobrowseconversationentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Cobrowseconversationentitylisting
type Cobrowseconversationentitylisting struct { 
    // Entities
    Entities []Cobrowseconversation `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Cobrowseconversationentitylisting) String() string {
     o.Entities = []Cobrowseconversation{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cobrowseconversationentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Cobrowseconversationentitylisting

    if CobrowseconversationentitylistingMarshalled {
        return []byte("{}"), nil
    }
    CobrowseconversationentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Cobrowseconversation `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Cobrowseconversation{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

