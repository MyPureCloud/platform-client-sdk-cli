package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediapolicyconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediapolicyconditionsDud struct { 
    


    


    


    


    


    


    


    

}

// Messagemediapolicyconditions
type Messagemediapolicyconditions struct { 
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


    // CustomerParticipation
    CustomerParticipation string `json:"customerParticipation"`

}

// String returns a JSON representation of the model
func (o *Messagemediapolicyconditions) String() string {
     o.ForUsers = []User{{}} 
     o.DateRanges = []string{""} 
     o.ForQueues = []Queue{{}} 
     o.WrapupCodes = []Wrapupcode{{}} 
     o.Languages = []Language{{}} 
    
     o.Teams = []Team{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediapolicyconditions) MarshalJSON() ([]byte, error) {
    type Alias Messagemediapolicyconditions

    if MessagemediapolicyconditionsMarshalled {
        return []byte("{}"), nil
    }
    MessagemediapolicyconditionsMarshalled = true

    return json.Marshal(&struct {
        
        ForUsers []User `json:"forUsers"`
        
        DateRanges []string `json:"dateRanges"`
        
        ForQueues []Queue `json:"forQueues"`
        
        WrapupCodes []Wrapupcode `json:"wrapupCodes"`
        
        Languages []Language `json:"languages"`
        
        TimeAllowed Timeallowed `json:"timeAllowed"`
        
        Teams []Team `json:"teams"`
        
        CustomerParticipation string `json:"customerParticipation"`
        *Alias
    }{

        
        ForUsers: []User{{}},
        


        
        DateRanges: []string{""},
        


        
        ForQueues: []Queue{{}},
        


        
        WrapupCodes: []Wrapupcode{{}},
        


        
        Languages: []Language{{}},
        


        


        
        Teams: []Team{{}},
        


        

        Alias: (*Alias)(u),
    })
}

