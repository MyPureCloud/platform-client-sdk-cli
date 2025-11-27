package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialnormalizedmessageentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialnormalizedmessageentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Opensocialnormalizedmessageentitylisting
type Opensocialnormalizedmessageentitylisting struct { 
    // Entities
    Entities []Opensocialmedianormalizedmessage `json:"entities"`


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
func (o *Opensocialnormalizedmessageentitylisting) String() string {
     o.Entities = []Opensocialmedianormalizedmessage{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialnormalizedmessageentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Opensocialnormalizedmessageentitylisting

    if OpensocialnormalizedmessageentitylistingMarshalled {
        return []byte("{}"), nil
    }
    OpensocialnormalizedmessageentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Opensocialmedianormalizedmessage `json:"entities"`
        
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

        
        Entities: []Opensocialmedianormalizedmessage{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

