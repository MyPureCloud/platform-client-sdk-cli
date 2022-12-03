package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityDud struct { 
    Id string `json:"id"`


    DateCompleted time.Time `json:"dateCompleted"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    PercentageScore float32 `json:"percentageScore"`


    IsPassed bool `json:"isPassed"`


    


    SelfUri string `json:"selfUri"`


    


    


    


    


    


    


    

}

// Developmentactivity - Development Activity object
type Developmentactivity struct { 
    


    


    


    


    


    


    // IsLatest - True if this is the latest version of assignment assigned to the user
    IsLatest bool `json:"isLatest"`


    


    // Name - The name of the activity
    Name string `json:"name"`


    // VarType - The type of activity
    VarType string `json:"type"`


    // Status - The status of the activity
    Status string `json:"status"`


    // DateDue - Due date for completion of the activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`


    // Facilitator - Facilitator of the activity
    Facilitator Userreference `json:"facilitator"`


    // Attendees - List of users attending the activity
    Attendees []Userreference `json:"attendees"`


    // IsOverdue - Indicates if the activity is overdue
    IsOverdue bool `json:"isOverdue"`

}

// String returns a JSON representation of the model
func (o *Developmentactivity) String() string {
    
    
    
    
    
    
     o.Attendees = []Userreference{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivity) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivity

    if DevelopmentactivityMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityMarshalled = true

    return json.Marshal(&struct {
        
        IsLatest bool `json:"isLatest"`
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Status string `json:"status"`
        
        DateDue time.Time `json:"dateDue"`
        
        Facilitator Userreference `json:"facilitator"`
        
        Attendees []Userreference `json:"attendees"`
        
        IsOverdue bool `json:"isOverdue"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Attendees: []Userreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

