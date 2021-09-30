package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentresponseDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Description string `json:"description"`


    DateStart time.Time `json:"dateStart"`


    LengthInMinutes int `json:"lengthInMinutes"`


    Status string `json:"status"`


    Facilitator Userreference `json:"facilitator"`


    Attendees []Userreference `json:"attendees"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    Conversations []Conversationreference `json:"conversations"`


    Documents []Documentreference `json:"documents"`


    IsOverdue bool `json:"isOverdue"`


    SelfUri string `json:"selfUri"`

}

// Coachingappointmentresponse - Coaching appointment response
type Coachingappointmentresponse struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Coachingappointmentresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentresponse) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentresponse

    if CoachingappointmentresponseMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentresponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

