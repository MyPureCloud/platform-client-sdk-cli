package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrenceresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanoccurrenceresponse
type Activityplanoccurrenceresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // RunDate - The date the associated activity plan was run, in ISO-8601 format
    RunDate time.Time `json:"runDate"`


    // SessionDateRange - The dates potentially spanned by the sessions in this occurrence, even if no sessions currently exist
    SessionDateRange Requireddaterange `json:"sessionDateRange"`


    // Facilitated - Whether the run sessions are facilitated
    Facilitated bool `json:"facilitated"`


    // Edited - Whether this occurrence is edited
    Edited bool `json:"edited"`


    // Violations - The violations for this occurrence
    Violations []Activityplanoccurrenceviolation `json:"violations"`


    // Sessions - The sessions scheduled by the run
    Sessions []Activityplanoccurrencesessionwithattendees `json:"sessions"`


    // UnscheduledAttendees - The attendees who could not be scheduled for this occurrence
    UnscheduledAttendees []Activityplanunscheduledattendee `json:"unscheduledAttendees"`


    // AffectedSchedules - Summaries of the schedules affected by this occurrence
    AffectedSchedules []Buschedulesummary `json:"affectedSchedules"`


    // CreatedBy - The user who started the job which created this occurrence
    CreatedBy Userreference `json:"createdBy"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrenceresponse) String() string {
    
    
    
    
    
     o.Violations = []Activityplanoccurrenceviolation{{}} 
     o.Sessions = []Activityplanoccurrencesessionwithattendees{{}} 
     o.UnscheduledAttendees = []Activityplanunscheduledattendee{{}} 
     o.AffectedSchedules = []Buschedulesummary{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrenceresponse

    if ActivityplanoccurrenceresponseMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        RunDate time.Time `json:"runDate"`
        
        SessionDateRange Requireddaterange `json:"sessionDateRange"`
        
        Facilitated bool `json:"facilitated"`
        
        Edited bool `json:"edited"`
        
        Violations []Activityplanoccurrenceviolation `json:"violations"`
        
        Sessions []Activityplanoccurrencesessionwithattendees `json:"sessions"`
        
        UnscheduledAttendees []Activityplanunscheduledattendee `json:"unscheduledAttendees"`
        
        AffectedSchedules []Buschedulesummary `json:"affectedSchedules"`
        
        CreatedBy Userreference `json:"createdBy"`
        *Alias
    }{

        


        


        


        


        


        
        Violations: []Activityplanoccurrenceviolation{{}},
        


        
        Sessions: []Activityplanoccurrencesessionwithattendees{{}},
        


        
        UnscheduledAttendees: []Activityplanunscheduledattendee{{}},
        


        
        AffectedSchedules: []Buschedulesummary{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

