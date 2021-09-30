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


    // Duration
    Duration Durationcondition `json:"duration"`

}

// String returns a JSON representation of the model
func (o *Chatmediapolicyconditions) String() string {
    
    
     o.ForUsers = []User{{}} 
    
    
    
     o.DateRanges = []string{""} 
    
    
    
     o.ForQueues = []Queue{{}} 
    
    
    
     o.WrapupCodes = []Wrapupcode{{}} 
    
    
    
     o.Languages = []Language{{}} 
    
    
    
    
    
    
    
    
    
    

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
        ForUsers []User `json:"forUsers"`
        
        DateRanges []string `json:"dateRanges"`
        
        ForQueues []Queue `json:"forQueues"`
        
        WrapupCodes []Wrapupcode `json:"wrapupCodes"`
        
        Languages []Language `json:"languages"`
        
        TimeAllowed Timeallowed `json:"timeAllowed"`
        
        Duration Durationcondition `json:"duration"`
        
        *Alias
    }{
        

        
        ForUsers: []User{{}},
        

        

        
        DateRanges: []string{""},
        

        

        
        ForQueues: []Queue{{}},
        

        

        
        WrapupCodes: []Wrapupcode{{}},
        

        

        
        Languages: []Language{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

