package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SettimeofflimitvaluesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SettimeofflimitvaluesrequestDud struct { 
    


    

}

// Settimeofflimitvaluesrequest - The non-empty list of the time off limit value intervals
type Settimeofflimitvaluesrequest struct { 
    // Values
    Values []Timeofflimitrange `json:"values"`


    // Metadata - Version metadata for the time off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Settimeofflimitvaluesrequest) String() string {
    
    
     o.Values = []Timeofflimitrange{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Settimeofflimitvaluesrequest) MarshalJSON() ([]byte, error) {
    type Alias Settimeofflimitvaluesrequest

    if SettimeofflimitvaluesrequestMarshalled {
        return []byte("{}"), nil
    }
    SettimeofflimitvaluesrequestMarshalled = true

    return json.Marshal(&struct { 
        Values []Timeofflimitrange `json:"values"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        
        Values: []Timeofflimitrange{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

