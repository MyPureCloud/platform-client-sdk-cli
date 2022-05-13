package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProvisioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProvisioninfoDud struct { 
    


    


    

}

// Provisioninfo
type Provisioninfo struct { 
    // Time - The time at which this phone was provisioned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`


    // Source - The source of the provisioning
    Source string `json:"source"`


    // ErrorInfo - The error information from the provision process, if any
    ErrorInfo string `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Provisioninfo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Provisioninfo) MarshalJSON() ([]byte, error) {
    type Alias Provisioninfo

    if ProvisioninfoMarshalled {
        return []byte("{}"), nil
    }
    ProvisioninfoMarshalled = true

    return json.Marshal(&struct {
        
        Time time.Time `json:"time"`
        
        Source string `json:"source"`
        
        ErrorInfo string `json:"errorInfo"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

