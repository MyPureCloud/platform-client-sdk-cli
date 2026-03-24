package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourceschedulesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourceschedulesettingsDud struct { 
    


    

}

// V3sourceschedulesettings
type V3sourceschedulesettings struct { 
    // DateStart - The date-time value of the first sync run. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // Period - The time between syncs, in hours, until a new sync is run.
    Period int `json:"period"`

}

// String returns a JSON representation of the model
func (o *V3sourceschedulesettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourceschedulesettings) MarshalJSON() ([]byte, error) {
    type Alias V3sourceschedulesettings

    if V3sourceschedulesettingsMarshalled {
        return []byte("{}"), nil
    }
    V3sourceschedulesettingsMarshalled = true

    return json.Marshal(&struct {
        
        DateStart time.Time `json:"dateStart"`
        
        Period int `json:"period"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

