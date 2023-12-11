package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LogcapturequeryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LogcapturequeryrequestDud struct { 
    


    


    

}

// Logcapturequeryrequest
type Logcapturequeryrequest struct { 
    // Interval - Date and time range to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // UserId - Id of the user to query.
    UserId string `json:"userId"`


    // SortOrder - Order of results. Default order is ASC.
    SortOrder string `json:"sortOrder"`

}

// String returns a JSON representation of the model
func (o *Logcapturequeryrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Logcapturequeryrequest) MarshalJSON() ([]byte, error) {
    type Alias Logcapturequeryrequest

    if LogcapturequeryrequestMarshalled {
        return []byte("{}"), nil
    }
    LogcapturequeryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        UserId string `json:"userId"`
        
        SortOrder string `json:"sortOrder"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

