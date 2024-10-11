package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialkeywordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialkeywordDud struct { 
    


    

}

// Socialkeyword
type Socialkeyword struct { 
    // Includes - List of keywords that must be included
    Includes []string `json:"includes"`


    // Excludes - List of keywords that must be excluded
    Excludes []string `json:"excludes"`

}

// String returns a JSON representation of the model
func (o *Socialkeyword) String() string {
     o.Includes = []string{""} 
     o.Excludes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialkeyword) MarshalJSON() ([]byte, error) {
    type Alias Socialkeyword

    if SocialkeywordMarshalled {
        return []byte("{}"), nil
    }
    SocialkeywordMarshalled = true

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

