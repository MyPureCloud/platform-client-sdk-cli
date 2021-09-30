package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssignedwrapupcodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssignedwrapupcodeDud struct { 
    


    


    


    


    

}

// Assignedwrapupcode
type Assignedwrapupcode struct { 
    // Code - The user configured wrap up code id.
    Code string `json:"code"`


    // Notes - Text entered by the agent to describe the call or disposition.
    Notes string `json:"notes"`


    // Tags - List of tags selected by the agent to describe the call or disposition.
    Tags []string `json:"tags"`


    // DurationSeconds - The duration in seconds of the wrap-up segment.
    DurationSeconds int `json:"durationSeconds"`


    // EndTime - The timestamp when the wrap-up segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`

}

// String returns a JSON representation of the model
func (o *Assignedwrapupcode) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Tags = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assignedwrapupcode) MarshalJSON() ([]byte, error) {
    type Alias Assignedwrapupcode

    if AssignedwrapupcodeMarshalled {
        return []byte("{}"), nil
    }
    AssignedwrapupcodeMarshalled = true

    return json.Marshal(&struct { 
        Code string `json:"code"`
        
        Notes string `json:"notes"`
        
        Tags []string `json:"tags"`
        
        DurationSeconds int `json:"durationSeconds"`
        
        EndTime time.Time `json:"endTime"`
        
        *Alias
    }{
        

        

        

        

        

        
        Tags: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

