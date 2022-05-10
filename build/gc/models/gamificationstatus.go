package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GamificationstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GamificationstatusDud struct { 
    


    


    


    

}

// Gamificationstatus
type Gamificationstatus struct { 
    // IsActive - Gamification status of the organization.
    IsActive bool `json:"isActive"`


    // DateStart - Gamification start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`


    // AutomaticUserAssignment - Automatic assignment of users to the default profile
    AutomaticUserAssignment bool `json:"automaticUserAssignment"`


    // DateStartPersonalBest - Personal best aggregation starting date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStartPersonalBest time.Time `json:"dateStartPersonalBest"`

}

// String returns a JSON representation of the model
func (o *Gamificationstatus) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Gamificationstatus) MarshalJSON() ([]byte, error) {
    type Alias Gamificationstatus

    if GamificationstatusMarshalled {
        return []byte("{}"), nil
    }
    GamificationstatusMarshalled = true

    return json.Marshal(&struct { 
        IsActive bool `json:"isActive"`
        
        DateStart time.Time `json:"dateStart"`
        
        AutomaticUserAssignment bool `json:"automaticUserAssignment"`
        
        DateStartPersonalBest time.Time `json:"dateStartPersonalBest"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

