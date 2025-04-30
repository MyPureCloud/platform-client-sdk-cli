package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingintegrationentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingintegrationentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Messagingintegrationentitylisting
type Messagingintegrationentitylisting struct { 
    // Entities
    Entities []Messagingintegration `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Messagingintegrationentitylisting) String() string {
     o.Entities = []Messagingintegration{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingintegrationentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Messagingintegrationentitylisting

    if MessagingintegrationentitylistingMarshalled {
        return []byte("{}"), nil
    }
    MessagingintegrationentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Messagingintegration `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Messagingintegration{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

