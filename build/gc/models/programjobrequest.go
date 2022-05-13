package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramjobrequestDud struct { 
    

}

// Programjobrequest
type Programjobrequest struct { 
    // ProgramIds - The ids of the programs used for this job
    ProgramIds []string `json:"programIds"`

}

// String returns a JSON representation of the model
func (o *Programjobrequest) String() string {
     o.ProgramIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Programjobrequest

    if ProgramjobrequestMarshalled {
        return []byte("{}"), nil
    }
    ProgramjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ProgramIds []string `json:"programIds"`
        *Alias
    }{

        
        ProgramIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

