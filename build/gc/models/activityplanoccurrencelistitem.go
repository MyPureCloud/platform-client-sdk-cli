package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrencelistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrencelistitemDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Activityplanoccurrencelistitem
type Activityplanoccurrencelistitem struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // RunDate - The date the associated activity plan was run, in ISO-8601 format
    RunDate time.Time `json:"runDate"`


    // SessionDateRange - The dates potentially spanned by the sessions in this occurrence, even if no sessions currently exist
    SessionDateRange Requireddaterange `json:"sessionDateRange"`


    // Facilitated - Whether the run sessions are facilitated
    Facilitated bool `json:"facilitated"`


    // SessionCount - The number of sessions that were scheduled
    SessionCount int `json:"sessionCount"`


    // ScheduledAttendeesCount - The number of attendees that were scheduled for this occurrence
    ScheduledAttendeesCount int `json:"scheduledAttendeesCount"`


    // UnscheduledAttendeesCount - The number of attendees that were unable to be scheduled for this occurrence
    UnscheduledAttendeesCount int `json:"unscheduledAttendeesCount"`


    // Edited - Whether this occurrence is edited
    Edited bool `json:"edited"`


    // ViolationCount - The number of violations for this occurrence
    ViolationCount int `json:"violationCount"`


    // AffectedScheduleIndexes - The index of the affected schedule summary in the schedules list
    AffectedScheduleIndexes []int `json:"affectedScheduleIndexes"`


    // CreatedBy - The user who started the job which created this occurrence. The id may be 'System' if it was an automated process
    CreatedBy Userreference `json:"createdBy"`


    

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencelistitem) String() string {
    
    
    
    
    
    
    
    
    
     o.AffectedScheduleIndexes = []int{0} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrencelistitem) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrencelistitem

    if ActivityplanoccurrencelistitemMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrencelistitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        RunDate time.Time `json:"runDate"`
        
        SessionDateRange Requireddaterange `json:"sessionDateRange"`
        
        Facilitated bool `json:"facilitated"`
        
        SessionCount int `json:"sessionCount"`
        
        ScheduledAttendeesCount int `json:"scheduledAttendeesCount"`
        
        UnscheduledAttendeesCount int `json:"unscheduledAttendeesCount"`
        
        Edited bool `json:"edited"`
        
        ViolationCount int `json:"violationCount"`
        
        AffectedScheduleIndexes []int `json:"affectedScheduleIndexes"`
        
        CreatedBy Userreference `json:"createdBy"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        AffectedScheduleIndexes: []int{0},
        


        


        

        Alias: (*Alias)(u),
    })
}

