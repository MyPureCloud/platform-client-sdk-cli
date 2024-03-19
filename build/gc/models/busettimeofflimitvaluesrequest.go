package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusettimeofflimitvaluesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusettimeofflimitvaluesrequestDud struct { 
    


    

}

// Busettimeofflimitvaluesrequest
type Busettimeofflimitvaluesrequest struct { 
    // Values
    Values []Butimeofflimitrange `json:"values"`


    // Metadata - Version metadata for the time-off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Busettimeofflimitvaluesrequest) String() string {
     o.Values = []Butimeofflimitrange{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Busettimeofflimitvaluesrequest) MarshalJSON() ([]byte, error) {
    type Alias Busettimeofflimitvaluesrequest

    if BusettimeofflimitvaluesrequestMarshalled {
        return []byte("{}"), nil
    }
    BusettimeofflimitvaluesrequestMarshalled = true

    return json.Marshal(&struct {
        
        Values []Butimeofflimitrange `json:"values"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        
        Values: []Butimeofflimitrange{{}},
        


        

        Alias: (*Alias)(u),
    })
}

