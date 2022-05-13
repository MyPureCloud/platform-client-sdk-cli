package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AftercallworkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AftercallworkDud struct { 
    


    


    

}

// Aftercallwork
type Aftercallwork struct { 
    // StartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // State
    State string `json:"state"`

}

// String returns a JSON representation of the model
func (o *Aftercallwork) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aftercallwork) MarshalJSON() ([]byte, error) {
    type Alias Aftercallwork

    if AftercallworkMarshalled {
        return []byte("{}"), nil
    }
    AftercallworkMarshalled = true

    return json.Marshal(&struct {
        
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

