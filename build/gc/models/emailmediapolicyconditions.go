package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailmediapolicyconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailmediapolicyconditionsDud struct { 
    


    


    


    


    


    

}

// Emailmediapolicyconditions
type Emailmediapolicyconditions struct { 
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

}

// String returns a JSON representation of the model
func (o *Emailmediapolicyconditions) String() string {
    
    
     o.ForUsers = []User{{}} 
    
    
    
     o.DateRanges = []string{""} 
    
    
    
     o.ForQueues = []Queue{{}} 
    
    
    
     o.WrapupCodes = []Wrapupcode{{}} 
    
    
    
     o.Languages = []Language{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailmediapolicyconditions) MarshalJSON() ([]byte, error) {
    type Alias Emailmediapolicyconditions

    if EmailmediapolicyconditionsMarshalled {
        return []byte("{}"), nil
    }
    EmailmediapolicyconditionsMarshalled = true

    return json.Marshal(&struct { 
        ForUsers []User `json:"forUsers"`
        
        DateRanges []string `json:"dateRanges"`
        
        ForQueues []Queue `json:"forQueues"`
        
        WrapupCodes []Wrapupcode `json:"wrapupCodes"`
        
        Languages []Language `json:"languages"`
        
        TimeAllowed Timeallowed `json:"timeAllowed"`
        
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

