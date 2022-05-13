package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatetimeofflimitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatetimeofflimitrequestDud struct { 
    


    

}

// Updatetimeofflimitrequest - Contains time off limit object property values to be updated.
type Updatetimeofflimitrequest struct { 
    // DefaultLimitMinutes - The default time off limit value in minutes per granularity
    DefaultLimitMinutes int `json:"defaultLimitMinutes"`


    // Metadata - Version metadata for the time off limit
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatetimeofflimitrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatetimeofflimitrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatetimeofflimitrequest

    if UpdatetimeofflimitrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatetimeofflimitrequestMarshalled = true

    return json.Marshal(&struct {
        
        DefaultLimitMinutes int `json:"defaultLimitMinutes"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

