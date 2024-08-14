package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewjoblistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewjoblistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Journeyviewjoblisting
type Journeyviewjoblisting struct { 
    // Entities
    Entities []Journeyviewjob `json:"entities"`


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
func (o *Journeyviewjoblisting) String() string {
     o.Entities = []Journeyviewjob{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewjoblisting) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewjoblisting

    if JourneyviewjoblistingMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewjoblistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Journeyviewjob `json:"entities"`
        
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

        
        Entities: []Journeyviewjob{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

