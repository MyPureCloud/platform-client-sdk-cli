package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingnotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingnotificationDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    


    ActionType string `json:"actionType"`


    Relationship string `json:"relationship"`


    DateStart time.Time `json:"dateStart"`


    LengthInMinutes int `json:"lengthInMinutes"`


    Status string `json:"status"`


    User Userreference `json:"user"`


    Appointment Coachingappointmentresponse `json:"appointment"`


    SelfUri string `json:"selfUri"`

}

// Coachingnotification
type Coachingnotification struct { 
    


    


    // MarkedAsRead - Indicates if notification is read or unread
    MarkedAsRead bool `json:"markedAsRead"`


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Coachingnotification) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingnotification) MarshalJSON() ([]byte, error) {
    type Alias Coachingnotification

    if CoachingnotificationMarshalled {
        return []byte("{}"), nil
    }
    CoachingnotificationMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        MarkedAsRead bool `json:"markedAsRead"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

