package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RegisterarchitectexportjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RegisterarchitectexportjobDud struct { 
    

}

// Registerarchitectexportjob
type Registerarchitectexportjob struct { 
    // Flows - A list of the flows to be exported.
    Flows []Exportdetails `json:"flows"`

}

// String returns a JSON representation of the model
func (o *Registerarchitectexportjob) String() string {
     o.Flows = []Exportdetails{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Registerarchitectexportjob) MarshalJSON() ([]byte, error) {
    type Alias Registerarchitectexportjob

    if RegisterarchitectexportjobMarshalled {
        return []byte("{}"), nil
    }
    RegisterarchitectexportjobMarshalled = true

    return json.Marshal(&struct {
        
        Flows []Exportdetails `json:"flows"`
        *Alias
    }{

        
        Flows: []Exportdetails{{}},
        

        Alias: (*Alias)(u),
    })
}

