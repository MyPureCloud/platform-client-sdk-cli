package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactcolumntimezoneMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactcolumntimezoneDud struct { 
    


    ColumnType string `json:"columnType"`

}

// Contactcolumntimezone
type Contactcolumntimezone struct { 
    // TimeZone - Time zone that the column matched to. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZone string `json:"timeZone"`


    

}

// String returns a JSON representation of the model
func (o *Contactcolumntimezone) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactcolumntimezone) MarshalJSON() ([]byte, error) {
    type Alias Contactcolumntimezone

    if ContactcolumntimezoneMarshalled {
        return []byte("{}"), nil
    }
    ContactcolumntimezoneMarshalled = true

    return json.Marshal(&struct {
        
        TimeZone string `json:"timeZone"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

