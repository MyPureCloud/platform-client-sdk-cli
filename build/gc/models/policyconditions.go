package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyconditionsDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Policyconditions
type Policyconditions struct { 
    // ForUsers - List of users to apply this policy to. Each user object can include the 'id' field containing the user's unique identifier. Example: [{\"id\":\"<userId>\"}].
    ForUsers []Policyuser `json:"forUsers"`


    // Directions
    Directions []string `json:"directions"`


    // DateRanges
    DateRanges []string `json:"dateRanges"`


    // MediaTypes
    MediaTypes []string `json:"mediaTypes"`


    // ForQueues
    ForQueues []Queue `json:"forQueues"`


    // Duration
    Duration Durationcondition `json:"duration"`


    // WrapupCodes
    WrapupCodes []Wrapupcode `json:"wrapupCodes"`


    // TimeAllowed
    TimeAllowed Timeallowed `json:"timeAllowed"`


    // Teams - Teams to match conversations against
    Teams []Team `json:"teams"`


    // CustomerParticipation - This condition is to filter out conversation with and without customer participation.
    CustomerParticipation string `json:"customerParticipation"`

}

// String returns a JSON representation of the model
func (o *Policyconditions) String() string {
     o.ForUsers = []Policyuser{{}} 
     o.Directions = []string{""} 
     o.DateRanges = []string{""} 
     o.MediaTypes = []string{""} 
     o.ForQueues = []Queue{{}} 
    
     o.WrapupCodes = []Wrapupcode{{}} 
    
     o.Teams = []Team{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyconditions) MarshalJSON() ([]byte, error) {
    type Alias Policyconditions

    if PolicyconditionsMarshalled {
        return []byte("{}"), nil
    }
    PolicyconditionsMarshalled = true

    return json.Marshal(&struct {
        
        ForUsers []Policyuser `json:"forUsers"`
        
        Directions []string `json:"directions"`
        
        DateRanges []string `json:"dateRanges"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        ForQueues []Queue `json:"forQueues"`
        
        Duration Durationcondition `json:"duration"`
        
        WrapupCodes []Wrapupcode `json:"wrapupCodes"`
        
        TimeAllowed Timeallowed `json:"timeAllowed"`
        
        Teams []Team `json:"teams"`
        
        CustomerParticipation string `json:"customerParticipation"`
        *Alias
    }{

        
        ForUsers: []Policyuser{{}},
        


        
        Directions: []string{""},
        


        
        DateRanges: []string{""},
        


        
        MediaTypes: []string{""},
        


        
        ForQueues: []Queue{{}},
        


        


        
        WrapupCodes: []Wrapupcode{{}},
        


        


        
        Teams: []Team{{}},
        


        

        Alias: (*Alias)(u),
    })
}

