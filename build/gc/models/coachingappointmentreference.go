package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Coachingappointmentreference
type Coachingappointmentreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Coachingappointmentreference) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentreference) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentreference

    if CoachingappointmentreferenceMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentreferenceMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

