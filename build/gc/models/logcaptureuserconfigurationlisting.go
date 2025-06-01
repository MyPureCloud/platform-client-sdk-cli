package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogcaptureuserconfigurationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogcaptureuserconfigurationlistingDud struct { 
    


    


    

}

// Logcaptureuserconfigurationlisting - List of log capture user configurations including total count and entities
type Logcaptureuserconfigurationlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Logcaptureuserconfiguration `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Logcaptureuserconfigurationlisting) String() string {
    
     o.Entities = []Logcaptureuserconfiguration{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logcaptureuserconfigurationlisting) MarshalJSON() ([]byte, error) {
    type Alias Logcaptureuserconfigurationlisting

    if LogcaptureuserconfigurationlistingMarshalled {
        return []byte("{}"), nil
    }
    LogcaptureuserconfigurationlistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Logcaptureuserconfiguration `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Logcaptureuserconfiguration{{}},
        


        

        Alias: (*Alias)(u),
    })
}

