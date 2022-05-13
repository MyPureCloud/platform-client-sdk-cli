package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationaddressverificationdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationaddressverificationdetailsDud struct { 
    


    


    


    

}

// Locationaddressverificationdetails
type Locationaddressverificationdetails struct { 
    // Status - Status of address verification process
    Status string `json:"status"`


    // DateFinished - Finished time of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateFinished time.Time `json:"dateFinished"`


    // DateStarted - Time started of address verification process. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // Service - Third party service used for address verification
    Service string `json:"service"`

}

// String returns a JSON representation of the model
func (o *Locationaddressverificationdetails) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationaddressverificationdetails) MarshalJSON() ([]byte, error) {
    type Alias Locationaddressverificationdetails

    if LocationaddressverificationdetailsMarshalled {
        return []byte("{}"), nil
    }
    LocationaddressverificationdetailsMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        DateFinished time.Time `json:"dateFinished"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        Service string `json:"service"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

