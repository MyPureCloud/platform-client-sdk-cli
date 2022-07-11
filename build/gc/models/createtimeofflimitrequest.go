package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatetimeofflimitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatetimeofflimitrequestDud struct { 
    


    

}

// Createtimeofflimitrequest
type Createtimeofflimitrequest struct { 
    // Granularity - Granularity choice for time off limit. If not specified, 'Daily' is assumed
    Granularity string `json:"granularity"`


    // DefaultLimitMinutes - The default limit value in minutes per granularity. If not specified, then 0 is assumed, which means there are no time off minutes available
    DefaultLimitMinutes int `json:"defaultLimitMinutes"`

}

// String returns a JSON representation of the model
func (o *Createtimeofflimitrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createtimeofflimitrequest) MarshalJSON() ([]byte, error) {
    type Alias Createtimeofflimitrequest

    if CreatetimeofflimitrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatetimeofflimitrequestMarshalled = true

    return json.Marshal(&struct {
        
        Granularity string `json:"granularity"`
        
        DefaultLimitMinutes int `json:"defaultLimitMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

