package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IvrMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IvrDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Ivr - Defines the phone numbers, operating hours, and the Architect flows to execute for an IVR.
type Ivr struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


    // Dnis - The phone number(s) to contact the IVR by.  Each phone number must be unique and not in use by another resource.  For example, a user and an iVR cannot have the same phone number.
    Dnis []string `json:"dnis"`


    // OpenHoursFlow - The Architect flow to execute during the hours an organization is open.
    OpenHoursFlow Domainentityref `json:"openHoursFlow"`


    // ClosedHoursFlow - The Architect flow to execute during the hours an organization is closed.
    ClosedHoursFlow Domainentityref `json:"closedHoursFlow"`


    // HolidayHoursFlow - The Architect flow to execute during an organization's holiday hours.
    HolidayHoursFlow Domainentityref `json:"holidayHoursFlow"`


    // ScheduleGroup - The schedule group defining the open and closed hours for an organization.  If this is provided, an open flow and a closed flow must be specified as well.
    ScheduleGroup Domainentityref `json:"scheduleGroup"`


    

}

// String returns a JSON representation of the model
func (o *Ivr) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Dnis = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ivr) MarshalJSON() ([]byte, error) {
    type Alias Ivr

    if IvrMarshalled {
        return []byte("{}"), nil
    }
    IvrMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        Dnis []string `json:"dnis"`
        
        OpenHoursFlow Domainentityref `json:"openHoursFlow"`
        
        ClosedHoursFlow Domainentityref `json:"closedHoursFlow"`
        
        HolidayHoursFlow Domainentityref `json:"holidayHoursFlow"`
        
        ScheduleGroup Domainentityref `json:"scheduleGroup"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Dnis: []string{""},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

