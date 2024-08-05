package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftbusettingsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftbusettingsresponseDud struct { 
    


    


    


    

}

// Alternativeshiftbusettingsresponse
type Alternativeshiftbusettingsresponse struct { 
    // EnabledGranularities - The granularity at which alternative shifts is allowed. An empty list means Alternative Shifts is disabled
    EnabledGranularities []string `json:"enabledGranularities"`


    // MinMinutesBeforeStartTime - The minimum number of minutes before the start of a shift that an alternative shift can be automatically approved
    MinMinutesBeforeStartTime int `json:"minMinutesBeforeStartTime"`


    // RetainedActivityCategories - Categories of activities that are required to remain at the same time slot for the alternative shifts offered. An empty list represents no retained activities
    RetainedActivityCategories []string `json:"retainedActivityCategories"`


    // Metadata - Version metadata for this business unit's alternative shift settings
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftbusettingsresponse) String() string {
     o.EnabledGranularities = []string{""} 
    
     o.RetainedActivityCategories = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftbusettingsresponse) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftbusettingsresponse

    if AlternativeshiftbusettingsresponseMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftbusettingsresponseMarshalled = true

    return json.Marshal(&struct {
        
        EnabledGranularities []string `json:"enabledGranularities"`
        
        MinMinutesBeforeStartTime int `json:"minMinutesBeforeStartTime"`
        
        RetainedActivityCategories []string `json:"retainedActivityCategories"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        
        EnabledGranularities: []string{""},
        


        


        
        RetainedActivityCategories: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

