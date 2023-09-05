package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimateavailablefulldaytimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimateavailablefulldaytimeoffrequestDud struct { 
    


    

}

// Estimateavailablefulldaytimeoffrequest
type Estimateavailablefulldaytimeoffrequest struct { 
    // Date - Date in yyyy-MM-dd format for full day request. Should be interpreted in the business unit's configured time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Date time.Time `json:"date"`


    // RequestedDurationMinutes - A requested length of time off request in minutes. If the value is null, then the system will use activity code length setting
    RequestedDurationMinutes int `json:"requestedDurationMinutes"`

}

// String returns a JSON representation of the model
func (o *Estimateavailablefulldaytimeoffrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimateavailablefulldaytimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Estimateavailablefulldaytimeoffrequest

    if EstimateavailablefulldaytimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    EstimateavailablefulldaytimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        Date time.Time `json:"date"`
        
        RequestedDurationMinutes int `json:"requestedDurationMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

