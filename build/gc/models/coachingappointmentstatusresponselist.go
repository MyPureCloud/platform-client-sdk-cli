package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentstatusresponselistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentstatusresponselistDud struct { 
    


    


    


    


    

}

// Coachingappointmentstatusresponselist
type Coachingappointmentstatusresponselist struct { 
    // Entities
    Entities []Coachingappointmentstatusresponse `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusresponselist) String() string {
    
    
     o.Entities = []Coachingappointmentstatusresponse{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentstatusresponselist) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentstatusresponselist

    if CoachingappointmentstatusresponselistMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentstatusresponselistMarshalled = true

    return json.Marshal(&struct { 
        Entities []Coachingappointmentstatusresponse `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Coachingappointmentstatusresponse{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

