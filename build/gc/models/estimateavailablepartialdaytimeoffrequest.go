package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimateavailablepartialdaytimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimateavailablepartialdaytimeoffrequestDud struct { 
    


    

}

// Estimateavailablepartialdaytimeoffrequest
type Estimateavailablepartialdaytimeoffrequest struct { 
    // Date - Start date-time in ISO-8601 format for partial day request
    Date time.Time `json:"date"`


    // RequestedDurationMinutes - A requested length of time off request in minutes. If the value is null, then the system will use activity code length setting
    RequestedDurationMinutes int `json:"requestedDurationMinutes"`

}

// String returns a JSON representation of the model
func (o *Estimateavailablepartialdaytimeoffrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimateavailablepartialdaytimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Estimateavailablepartialdaytimeoffrequest

    if EstimateavailablepartialdaytimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    EstimateavailablepartialdaytimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        Date time.Time `json:"date"`
        
        RequestedDurationMinutes int `json:"requestedDurationMinutes"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

