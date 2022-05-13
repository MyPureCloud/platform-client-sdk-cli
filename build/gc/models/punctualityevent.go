package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PunctualityeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PunctualityeventDud struct { 
    DateScheduleStart time.Time `json:"dateScheduleStart"`


    DateStart time.Time `json:"dateStart"`


    LengthMinutes int `json:"lengthMinutes"`


    Description string `json:"description"`


    ActivityCodeId string `json:"activityCodeId"`


    ActivityCode string `json:"activityCode"`


    Category string `json:"category"`


    Points int `json:"points"`


    Delta float64 `json:"delta"`


    

}

// Punctualityevent
type Punctualityevent struct { 
    


    


    


    


    


    


    


    


    


    // Bullseye
    Bullseye bool `json:"bullseye"`

}

// String returns a JSON representation of the model
func (o *Punctualityevent) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Punctualityevent) MarshalJSON() ([]byte, error) {
    type Alias Punctualityevent

    if PunctualityeventMarshalled {
        return []byte("{}"), nil
    }
    PunctualityeventMarshalled = true

    return json.Marshal(&struct {
        
        Bullseye bool `json:"bullseye"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

