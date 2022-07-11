package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeofflimitMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeofflimitDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Timeofflimit
type Timeofflimit struct { 
    


    // Granularity - Granularity choice for the time off limit
    Granularity string `json:"granularity"`


    // DefaultLimitMinutes - The default time off limit value in minutes per granularity interval
    DefaultLimitMinutes int `json:"defaultLimitMinutes"`


    // Metadata - Version metadata for the time off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Timeofflimit) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeofflimit) MarshalJSON() ([]byte, error) {
    type Alias Timeofflimit

    if TimeofflimitMarshalled {
        return []byte("{}"), nil
    }
    TimeofflimitMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        
        DefaultLimitMinutes int `json:"defaultLimitMinutes"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

