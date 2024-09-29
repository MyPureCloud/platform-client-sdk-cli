package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourcelastsyncMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourcelastsyncDud struct { 
    


    


    


    

}

// Sourcelastsync
type Sourcelastsync struct { 
    // State - State of the last synchronization.
    State string `json:"state"`


    // DateStarted - Last synchronization start-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // DateEnded - Last synchronization end-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnded time.Time `json:"dateEnded"`


    // VarError - Optional error message of the last synchronization.
    VarError Errorbody `json:"error"`

}

// String returns a JSON representation of the model
func (o *Sourcelastsync) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourcelastsync) MarshalJSON() ([]byte, error) {
    type Alias Sourcelastsync

    if SourcelastsyncMarshalled {
        return []byte("{}"), nil
    }
    SourcelastsyncMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        DateEnded time.Time `json:"dateEnded"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

