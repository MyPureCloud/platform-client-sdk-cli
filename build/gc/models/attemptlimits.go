package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AttemptlimitsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AttemptlimitsDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Attemptlimits
type Attemptlimits struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // MaxAttemptsPerContact - The maximum number of times a contact can be called within the resetPeriod. Required if maxAttemptsPerNumber is not defined.
    MaxAttemptsPerContact int `json:"maxAttemptsPerContact"`


    // MaxAttemptsPerNumber - The maximum number of times a phone number can be called within the resetPeriod. Required if maxAttemptsPerContact is not defined.
    MaxAttemptsPerNumber int `json:"maxAttemptsPerNumber"`


    // TimeZoneId - If the resetPeriod is TODAY, this specifies the timezone in which TODAY occurs. Required if the resetPeriod is TODAY.
    TimeZoneId string `json:"timeZoneId"`


    // ResetPeriod - After how long the number of attempts will be set back to 0. Defaults to NEVER.
    ResetPeriod string `json:"resetPeriod"`


    // RecallEntries - Configuration for recall attempts.
    RecallEntries map[string]Recallentry `json:"recallEntries"`


    

}

// String returns a JSON representation of the model
func (o *Attemptlimits) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RecallEntries = map[string]Recallentry{"": {}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Attemptlimits) MarshalJSON() ([]byte, error) {
    type Alias Attemptlimits

    if AttemptlimitsMarshalled {
        return []byte("{}"), nil
    }
    AttemptlimitsMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        
        
        Version int `json:"version"`
        
        MaxAttemptsPerContact int `json:"maxAttemptsPerContact"`
        
        MaxAttemptsPerNumber int `json:"maxAttemptsPerNumber"`
        
        TimeZoneId string `json:"timeZoneId"`
        
        ResetPeriod string `json:"resetPeriod"`
        
        RecallEntries map[string]Recallentry `json:"recallEntries"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        RecallEntries: map[string]Recallentry{"": {}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

