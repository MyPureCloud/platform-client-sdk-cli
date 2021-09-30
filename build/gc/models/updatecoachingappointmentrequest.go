package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatecoachingappointmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatecoachingappointmentrequestDud struct { 
    


    


    


    


    


    


    

}

// Updatecoachingappointmentrequest - Update coaching appointment request
type Updatecoachingappointmentrequest struct { 
    // Name - The name of coaching appointment.
    Name string `json:"name"`


    // Description - The description of coaching appointment.
    Description string `json:"description"`


    // DateStart - The date/time the coaching appointment starts. Times will be rounded down to the minute. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // LengthInMinutes - The duration of coaching appointment in minutes.
    LengthInMinutes int `json:"lengthInMinutes"`


    // ConversationIds - IDs of conversations associated with this coaching appointment.
    ConversationIds []string `json:"conversationIds"`


    // DocumentIds - IDs of documents associated with this coaching appointment.
    DocumentIds []string `json:"documentIds"`


    // Status - The status of the coaching appointment.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Updatecoachingappointmentrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ConversationIds = []string{""} 
    
    
    
     o.DocumentIds = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatecoachingappointmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatecoachingappointmentrequest

    if UpdatecoachingappointmentrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatecoachingappointmentrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateStart time.Time `json:"dateStart"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        ConversationIds []string `json:"conversationIds"`
        
        DocumentIds []string `json:"documentIds"`
        
        Status string `json:"status"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        ConversationIds: []string{""},
        

        

        
        DocumentIds: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

