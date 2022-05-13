package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimezoneentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimezoneentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Timezoneentitylisting
type Timezoneentitylisting struct { 
    // Entities
    Entities []Regiontimezone `json:"entities"`


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


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Timezoneentitylisting) String() string {
     o.Entities = []Regiontimezone{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timezoneentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Timezoneentitylisting

    if TimezoneentitylistingMarshalled {
        return []byte("{}"), nil
    }
    TimezoneentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Regiontimezone `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Regiontimezone{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

