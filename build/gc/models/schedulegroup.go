package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegroupDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Schedulegroup - A group of schedules that define the operating hours of an organization.
type Schedulegroup struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // TimeZone - The timezone the schedules are a part of.  This is not a schedule property to allow a schedule to be used in multiple timezones.
    TimeZone string `json:"timeZone"`


    // OpenSchedules - The schedules defining the hours an organization is open.
    OpenSchedules []Domainentityref `json:"openSchedules"`


    // ClosedSchedules - The schedules defining the hours an organization is closed.
    ClosedSchedules []Domainentityref `json:"closedSchedules"`


    // HolidaySchedules - The schedules defining the hours an organization is closed for the holidays.
    HolidaySchedules []Domainentityref `json:"holidaySchedules"`


    

}

// String returns a JSON representation of the model
func (o *Schedulegroup) String() string {
    
    
    
    
    
     o.OpenSchedules = []Domainentityref{{}} 
     o.ClosedSchedules = []Domainentityref{{}} 
     o.HolidaySchedules = []Domainentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegroup) MarshalJSON() ([]byte, error) {
    type Alias Schedulegroup

    if SchedulegroupMarshalled {
        return []byte("{}"), nil
    }
    SchedulegroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        TimeZone string `json:"timeZone"`
        
        OpenSchedules []Domainentityref `json:"openSchedules"`
        
        ClosedSchedules []Domainentityref `json:"closedSchedules"`
        
        HolidaySchedules []Domainentityref `json:"holidaySchedules"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        OpenSchedules: []Domainentityref{{}},
        


        
        ClosedSchedules: []Domainentityref{{}},
        


        
        HolidaySchedules: []Domainentityref{{}},
        


        

        Alias: (*Alias)(u),
    })
}

