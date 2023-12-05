package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallmediapolicyconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallmediapolicyconditionsDud struct { 
    


    


    


    


    


    


    


    


    

}

// Callmediapolicyconditions
type Callmediapolicyconditions struct { 
    // ForUsers
    ForUsers []User `json:"forUsers"`


    // DateRanges
    DateRanges []string `json:"dateRanges"`


    // ForQueues
    ForQueues []Queue `json:"forQueues"`


    // WrapupCodes
    WrapupCodes []Wrapupcode `json:"wrapupCodes"`


    // Languages
    Languages []Language `json:"languages"`


    // TimeAllowed
    TimeAllowed Timeallowed `json:"timeAllowed"`


    // Teams - Teams to match conversations against
    Teams []Team `json:"teams"`


    // Directions
    Directions []string `json:"directions"`


    // Duration
    Duration Durationcondition `json:"duration"`

}

// String returns a JSON representation of the model
func (o *Callmediapolicyconditions) String() string {
     o.ForUsers = []User{{}} 
     o.DateRanges = []string{""} 
     o.ForQueues = []Queue{{}} 
     o.WrapupCodes = []Wrapupcode{{}} 
     o.Languages = []Language{{}} 
    
     o.Teams = []Team{{}} 
     o.Directions = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callmediapolicyconditions) MarshalJSON() ([]byte, error) {
    type Alias Callmediapolicyconditions

    if CallmediapolicyconditionsMarshalled {
        return []byte("{}"), nil
    }
    CallmediapolicyconditionsMarshalled = true

    return json.Marshal(&struct {
        
        ForUsers []User `json:"forUsers"`
        
        DateRanges []string `json:"dateRanges"`
        
        ForQueues []Queue `json:"forQueues"`
        
        WrapupCodes []Wrapupcode `json:"wrapupCodes"`
        
        Languages []Language `json:"languages"`
        
        TimeAllowed Timeallowed `json:"timeAllowed"`
        
        Teams []Team `json:"teams"`
        
        Directions []string `json:"directions"`
        
        Duration Durationcondition `json:"duration"`
        *Alias
    }{

        
        ForUsers: []User{{}},
        


        
        DateRanges: []string{""},
        


        
        ForQueues: []Queue{{}},
        


        
        WrapupCodes: []Wrapupcode{{}},
        


        
        Languages: []Language{{}},
        


        


        
        Teams: []Team{{}},
        


        
        Directions: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

