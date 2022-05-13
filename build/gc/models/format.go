package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormatMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormatDud struct { 
    

}

// Format
type Format struct { 
    // Flags - The Set of prompt segment format flags i.e. each entry is a part of describing the overall format. E.g. \"format\": { \"flags\": [StringPlayChars] }
    Flags []string `json:"flags"`

}

// String returns a JSON representation of the model
func (o *Format) String() string {
     o.Flags = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Format) MarshalJSON() ([]byte, error) {
    type Alias Format

    if FormatMarshalled {
        return []byte("{}"), nil
    }
    FormatMarshalled = true

    return json.Marshal(&struct {
        
        Flags []string `json:"flags"`
        *Alias
    }{

        
        Flags: []string{""},
        

        Alias: (*Alias)(u),
    })
}

