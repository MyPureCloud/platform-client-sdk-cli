package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatmediapolicyconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatmediapolicyconditionsDud struct { 
    


    


    


    


    


    


    


    

}

// Chatmediapolicyconditions
type Chatmediapolicyconditions struct { 
    // ForUsers - List of users to apply this policy to. Each user object can include the 'id' field containing the user's unique identifier. Example: [{\"id\":\"<userId>\"}].
    ForUsers []Policyuser `json:"forUsers"`


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


    // Duration
    Duration Durationcondition `json:"duration"`

}

// String returns a JSON representation of the model
func (o *Chatmediapolicyconditions) String() string {
     o.ForUsers = []Policyuser{{}} 
     o.DateRanges = []string{""} 
     o.ForQueues = []Queue{{}} 
     o.WrapupCodes = []Wrapupcode{{}} 
     o.Languages = []Language{{}} 
    
     o.Teams = []Team{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatmediapolicyconditions) MarshalJSON() ([]byte, error) {
    type Alias Chatmediapolicyconditions

    if ChatmediapolicyconditionsMarshalled {
        return []byte("{}"), nil
    }
    ChatmediapolicyconditionsMarshalled = true

    return json.Marshal(&struct {
        
        ForUsers []Policyuser `json:"forUsers"`
        
        DateRanges []string `json:"dateRanges"`
        
        ForQueues []Queue `json:"forQueues"`
        
        WrapupCodes []Wrapupcode `json:"wrapupCodes"`
        
        Languages []Language `json:"languages"`
        
        TimeAllowed Timeallowed `json:"timeAllowed"`
        
        Teams []Team `json:"teams"`
        
        Duration Durationcondition `json:"duration"`
        *Alias
    }{

        
        ForUsers: []Policyuser{{}},
        


        
        DateRanges: []string{""},
        


        
        ForQueues: []Queue{{}},
        


        
        WrapupCodes: []Wrapupcode{{}},
        


        
        Languages: []Language{{}},
        


        


        
        Teams: []Team{{}},
        


        

        Alias: (*Alias)(u),
    })
}

