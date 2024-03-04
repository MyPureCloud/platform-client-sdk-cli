package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowloglevelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowloglevelDud struct { 
    


    

}

// Flowloglevel - This is a table of settings per a loglevel that define what will be logged in executionData when enabled (true)
type Flowloglevel struct { 
    // Level - The logLevel for this characteristics set
    Level string `json:"level"`


    // Characteristics - Shows what characteristics are enabled for this log level
    Characteristics Flowcharacteristics `json:"characteristics"`

}

// String returns a JSON representation of the model
func (o *Flowloglevel) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowloglevel) MarshalJSON() ([]byte, error) {
    type Alias Flowloglevel

    if FlowloglevelMarshalled {
        return []byte("{}"), nil
    }
    FlowloglevelMarshalled = true

    return json.Marshal(&struct {
        
        Level string `json:"level"`
        
        Characteristics Flowcharacteristics `json:"characteristics"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

