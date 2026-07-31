package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeofflimitvaluesforgranularityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeofflimitvaluesforgranularityresponseDud struct { 
    


    


    


    

}

// Butimeofflimitvaluesforgranularityresponse
type Butimeofflimitvaluesforgranularityresponse struct { 
    // TimeOffLimit - The ID of the time-off limit
    TimeOffLimit Butimeofflimitreference `json:"timeOffLimit"`


    // Granularity - Granularity choice for time-off limit
    Granularity string `json:"granularity"`


    // LimitValues - Values for time-off limit
    LimitValues []Butimeofflimitvalues `json:"limitValues"`


    // Metadata - Version metadata for the time-off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Butimeofflimitvaluesforgranularityresponse) String() string {
    
    
     o.LimitValues = []Butimeofflimitvalues{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeofflimitvaluesforgranularityresponse) MarshalJSON() ([]byte, error) {
    type Alias Butimeofflimitvaluesforgranularityresponse

    if ButimeofflimitvaluesforgranularityresponseMarshalled {
        return []byte("{}"), nil
    }
    ButimeofflimitvaluesforgranularityresponseMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffLimit Butimeofflimitreference `json:"timeOffLimit"`
        
        Granularity string `json:"granularity"`
        
        LimitValues []Butimeofflimitvalues `json:"limitValues"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        
        LimitValues: []Butimeofflimitvalues{{}},
        


        

        Alias: (*Alias)(u),
    })
}

