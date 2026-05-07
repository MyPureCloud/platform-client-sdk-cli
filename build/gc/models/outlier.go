package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutlierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutlierDud struct { 
    


    


    

}

// Outlier
type Outlier struct { 
    // Timestamp - Timestamp of the outlier. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // Anomaly - The type of the anomaly
    Anomaly string `json:"anomaly"`


    // NormalizedValue - Normalized value of the outlier
    NormalizedValue float64 `json:"normalizedValue"`

}

// String returns a JSON representation of the model
func (o *Outlier) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outlier) MarshalJSON() ([]byte, error) {
    type Alias Outlier

    if OutlierMarshalled {
        return []byte("{}"), nil
    }
    OutlierMarshalled = true

    return json.Marshal(&struct {
        
        Timestamp time.Time `json:"timestamp"`
        
        Anomaly string `json:"anomaly"`
        
        NormalizedValue float64 `json:"normalizedValue"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

