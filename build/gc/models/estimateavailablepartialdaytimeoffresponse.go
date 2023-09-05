package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimateavailablepartialdaytimeoffresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimateavailablepartialdaytimeoffresponseDud struct { 
    


    


    


    

}

// Estimateavailablepartialdaytimeoffresponse
type Estimateavailablepartialdaytimeoffresponse struct { 
    // Date - Start date-time in ISO-8601 format for partial day request
    Date time.Time `json:"date"`


    // DurationMinutes - An estimation of time off request length in minutes
    DurationMinutes int `json:"durationMinutes"`


    // PayableMinutes - An estimation of payable part of time off request in minutes
    PayableMinutes int `json:"payableMinutes"`


    // Flexible - Whether there is flexibility for a user to choose different hours than the system estimated
    Flexible bool `json:"flexible"`

}

// String returns a JSON representation of the model
func (o *Estimateavailablepartialdaytimeoffresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimateavailablepartialdaytimeoffresponse) MarshalJSON() ([]byte, error) {
    type Alias Estimateavailablepartialdaytimeoffresponse

    if EstimateavailablepartialdaytimeoffresponseMarshalled {
        return []byte("{}"), nil
    }
    EstimateavailablepartialdaytimeoffresponseMarshalled = true

    return json.Marshal(&struct {
        
        Date time.Time `json:"date"`
        
        DurationMinutes int `json:"durationMinutes"`
        
        PayableMinutes int `json:"payableMinutes"`
        
        Flexible bool `json:"flexible"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

