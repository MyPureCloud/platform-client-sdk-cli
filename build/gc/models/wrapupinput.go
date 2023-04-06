package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WrapupinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WrapupinputDud struct { 
    


    


    


    


    


    


    


    

}

// Wrapupinput
type Wrapupinput struct { 
    // Code - The user configured wrap up code id.
    Code string `json:"code"`


    // Name - The user configured wrap up code name.
    Name string `json:"name"`


    // Notes - Text entered by the agent to describe the call or disposition.
    Notes string `json:"notes"`


    // Tags - List of tags selected by the agent to describe the call or disposition.
    Tags []string `json:"tags"`


    // DurationSeconds - The length of time in seconds that the agent spent doing after call work.
    DurationSeconds int `json:"durationSeconds"`


    // EndTime - The timestamp when the wrapup was finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // Provisional - Indicates if this is a pending save and should not require a code to be specified.  This allows someone to save some temporary wrapup that will be used later.
    Provisional bool `json:"provisional"`


    // DisableEndTimeUpdates - Prevent updates to wrapup end time when set to true.
    DisableEndTimeUpdates bool `json:"disableEndTimeUpdates"`

}

// String returns a JSON representation of the model
func (o *Wrapupinput) String() string {
    
    
    
     o.Tags = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wrapupinput) MarshalJSON() ([]byte, error) {
    type Alias Wrapupinput

    if WrapupinputMarshalled {
        return []byte("{}"), nil
    }
    WrapupinputMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Name string `json:"name"`
        
        Notes string `json:"notes"`
        
        Tags []string `json:"tags"`
        
        DurationSeconds int `json:"durationSeconds"`
        
        EndTime time.Time `json:"endTime"`
        
        Provisional bool `json:"provisional"`
        
        DisableEndTimeUpdates bool `json:"disableEndTimeUpdates"`
        *Alias
    }{

        


        


        


        
        Tags: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

