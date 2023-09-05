package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimateavailablefulldaytimeoffresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimateavailablefulldaytimeoffresponseDud struct { 
    


    


    


    

}

// Estimateavailablefulldaytimeoffresponse
type Estimateavailablefulldaytimeoffresponse struct { 
    // Date - Date in yyyy-MM-dd format for full day request. Should be interpreted in the business unit's configured time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Date time.Time `json:"date"`


    // DurationMinutes - An estimation of time off request length in minutes
    DurationMinutes int `json:"durationMinutes"`


    // PayableMinutes - An estimation of payable part of time off request in minutes
    PayableMinutes int `json:"payableMinutes"`


    // Flexible - Whether there is flexibility for a user to choose different hours than the system estimated
    Flexible bool `json:"flexible"`

}

// String returns a JSON representation of the model
func (o *Estimateavailablefulldaytimeoffresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimateavailablefulldaytimeoffresponse) MarshalJSON() ([]byte, error) {
    type Alias Estimateavailablefulldaytimeoffresponse

    if EstimateavailablefulldaytimeoffresponseMarshalled {
        return []byte("{}"), nil
    }
    EstimateavailablefulldaytimeoffresponseMarshalled = true

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

