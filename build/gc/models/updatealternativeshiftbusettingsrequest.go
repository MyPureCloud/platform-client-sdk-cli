package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatealternativeshiftbusettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatealternativeshiftbusettingsrequestDud struct { 
    


    


    


    

}

// Updatealternativeshiftbusettingsrequest
type Updatealternativeshiftbusettingsrequest struct { 
    // EnabledGranularities - The granularity at which alternative shifts is allowed. An empty list as the wrapped value will indicate alternative shifts is disabled
    EnabledGranularities Listwrapperalternativeshiftbusettingsgranularity `json:"enabledGranularities"`


    // MinMinutesBeforeStartTime - The minimum number of minutes before the start of a shift that an alternative shift can be automatically approved
    MinMinutesBeforeStartTime int `json:"minMinutesBeforeStartTime"`


    // RetainedActivityCategories - Categories of activities that are required to remain at the same time slot for the alternative shifts offered. An empty list indicates no retained activities
    RetainedActivityCategories Listwrapperalternativeshiftbusettingsactivitycategory `json:"retainedActivityCategories"`


    // Metadata - Version metadata for this business unit's alternative shift settings
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatealternativeshiftbusettingsrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatealternativeshiftbusettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatealternativeshiftbusettingsrequest

    if UpdatealternativeshiftbusettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatealternativeshiftbusettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        EnabledGranularities Listwrapperalternativeshiftbusettingsgranularity `json:"enabledGranularities"`
        
        MinMinutesBeforeStartTime int `json:"minMinutesBeforeStartTime"`
        
        RetainedActivityCategories Listwrapperalternativeshiftbusettingsactivitycategory `json:"retainedActivityCategories"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

