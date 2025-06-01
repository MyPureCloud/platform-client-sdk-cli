package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentresponselistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentresponselistDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Coachingappointmentresponselist
type Coachingappointmentresponselist struct { 
    // Entities
    Entities []Coachingappointmentresponse `json:"entities"`


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


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentresponselist) String() string {
     o.Entities = []Coachingappointmentresponse{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentresponselist) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentresponselist

    if CoachingappointmentresponselistMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentresponselistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Coachingappointmentresponse `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        LastUri string `json:"lastUri"`
        
        SelfUri string `json:"selfUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Coachingappointmentresponse{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

