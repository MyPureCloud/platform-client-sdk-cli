package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LicensebatchassignmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LicensebatchassignmentrequestDud struct { 
    

}

// Licensebatchassignmentrequest
type Licensebatchassignmentrequest struct { 
    // Assignments - The list of license assignment updates to make.
    Assignments []Licenseassignmentrequest `json:"assignments"`

}

// String returns a JSON representation of the model
func (o *Licensebatchassignmentrequest) String() string {
     o.Assignments = []Licenseassignmentrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Licensebatchassignmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Licensebatchassignmentrequest

    if LicensebatchassignmentrequestMarshalled {
        return []byte("{}"), nil
    }
    LicensebatchassignmentrequestMarshalled = true

    return json.Marshal(&struct {
        
        Assignments []Licenseassignmentrequest `json:"assignments"`
        *Alias
    }{

        
        Assignments: []Licenseassignmentrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

