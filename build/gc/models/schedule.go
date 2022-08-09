package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScheduleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScheduleDud struct { 
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

// Schedule - Defines a period of time to perform a specific action.  Each schedule must be associated with one or more schedule groups to be used.
type Schedule struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // Start - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
    Start time.Time `json:"start"`


    // End - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
    End time.Time `json:"end"`


    // Rrule - An iCal Recurrence Rule (RRULE) string. It is required to be set for schedules determining when upgrades to the Edge software can be applied.
    Rrule string `json:"rrule"`


    

}

// String returns a JSON representation of the model
func (o *Schedule) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedule) MarshalJSON() ([]byte, error) {
    type Alias Schedule

    if ScheduleMarshalled {
        return []byte("{}"), nil
    }
    ScheduleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Start time.Time `json:"start"`
        
        End time.Time `json:"end"`
        
        Rrule string `json:"rrule"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

