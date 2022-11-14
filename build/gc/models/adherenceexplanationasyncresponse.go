package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationasyncresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationasyncresponseDud struct { 
    

}

// Adherenceexplanationasyncresponse
type Adherenceexplanationasyncresponse struct { 
    // Job - A reference to the job that was started by the request
    Job Adherenceexplanationjobreference `json:"job"`

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationasyncresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationasyncresponse) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationasyncresponse

    if AdherenceexplanationasyncresponseMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationasyncresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Adherenceexplanationjobreference `json:"job"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

