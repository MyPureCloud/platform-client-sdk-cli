package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatecoachingappointmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatecoachingappointmentrequestDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Createcoachingappointmentrequest - Create coaching appointment request
type Createcoachingappointmentrequest struct { 
    // Name - The name of coaching appointment.
    Name string `json:"name"`


    // Description - The description of coaching appointment.
    Description string `json:"description"`


    // DateStart - The date/time the coaching appointment starts. Times will be rounded down to the minute. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // LengthInMinutes - The duration of coaching appointment in minutes.
    LengthInMinutes int `json:"lengthInMinutes"`


    // FacilitatorId - The facilitator ID of coaching appointment.
    FacilitatorId string `json:"facilitatorId"`


    // AttendeeIds - IDs of attendees in the coaching appointment.
    AttendeeIds []string `json:"attendeeIds"`


    // ConversationIds - IDs of conversations associated with this coaching appointment.
    ConversationIds []string `json:"conversationIds"`


    // DocumentIds - IDs of documents associated with this coaching appointment.
    DocumentIds []string `json:"documentIds"`


    // WfmSchedule - The Workforce Management schedule the appointment is associated with.
    WfmSchedule Wfmschedulereference `json:"wfmSchedule"`


    // ExternalLinks - The list of external links related to the appointment
    ExternalLinks []string `json:"externalLinks"`

}

// String returns a JSON representation of the model
func (o *Createcoachingappointmentrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.AttendeeIds = []string{""} 
    
    
    
     o.ConversationIds = []string{""} 
    
    
    
     o.DocumentIds = []string{""} 
    
    
    
    
    
    
    
     o.ExternalLinks = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createcoachingappointmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Createcoachingappointmentrequest

    if CreatecoachingappointmentrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatecoachingappointmentrequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateStart time.Time `json:"dateStart"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        FacilitatorId string `json:"facilitatorId"`
        
        AttendeeIds []string `json:"attendeeIds"`
        
        ConversationIds []string `json:"conversationIds"`
        
        DocumentIds []string `json:"documentIds"`
        
        WfmSchedule Wfmschedulereference `json:"wfmSchedule"`
        
        ExternalLinks []string `json:"externalLinks"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        AttendeeIds: []string{""},
        

        

        
        ConversationIds: []string{""},
        

        

        
        DocumentIds: []string{""},
        

        

        

        

        
        ExternalLinks: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

