package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodecontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodecontainerDud struct { 
    


    

}

// Activitycodecontainer - Container for a map of ActivityCodeId to ActivityCode
type Activitycodecontainer struct { 
    // ActivityCodes - Map of activity code id to activity code
    ActivityCodes map[string]Activitycode `json:"activityCodes"`


    // Metadata - Version metadata for the associated management unit's list of activity codes
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Activitycodecontainer) String() string {
     o.ActivityCodes = map[string]Activitycode{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycodecontainer) MarshalJSON() ([]byte, error) {
    type Alias Activitycodecontainer

    if ActivitycodecontainerMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodecontainerMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodes map[string]Activitycode `json:"activityCodes"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        
        ActivityCodes: map[string]Activitycode{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

