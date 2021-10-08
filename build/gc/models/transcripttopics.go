package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TranscripttopicsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TranscripttopicsDud struct { 
    


    

}

// Transcripttopics
type Transcripttopics struct { 
    // Includes - List of topics which need to be included in exact match criteria. This field is not mutually exclusive with excludes topic list.
    Includes []string `json:"includes"`


    // Excludes - List of topics which need to be excluded in exact match criteria. This field is not mutually exclusive with includes topic list.
    Excludes []string `json:"excludes"`

}

// String returns a JSON representation of the model
func (o *Transcripttopics) String() string {
    
    
     o.Includes = []string{""} 
    
    
    
     o.Excludes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transcripttopics) MarshalJSON() ([]byte, error) {
    type Alias Transcripttopics

    if TranscripttopicsMarshalled {
        return []byte("{}"), nil
    }
    TranscripttopicsMarshalled = true

    return json.Marshal(&struct { 
        Includes []string `json:"includes"`
        
        Excludes []string `json:"excludes"`
        
        *Alias
    }{
        

        
        Includes: []string{""},
        

        

        
        Excludes: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

