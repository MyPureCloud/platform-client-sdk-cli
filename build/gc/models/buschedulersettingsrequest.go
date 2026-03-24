package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuschedulersettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuschedulersettingsrequestDud struct { 
    


    

}

// Buschedulersettingsrequest
type Buschedulersettingsrequest struct { 
    // ConsistentServiceLevelSmoothing - Indicates whether to provide consistent service level smoothing in schedule generation for this business unit
    ConsistentServiceLevelSmoothing bool `json:"consistentServiceLevelSmoothing"`


    // Metadata - Version metadata for this business unit's scheduler settings
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Buschedulersettingsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buschedulersettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Buschedulersettingsrequest

    if BuschedulersettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    BuschedulersettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        ConsistentServiceLevelSmoothing bool `json:"consistentServiceLevelSmoothing"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

