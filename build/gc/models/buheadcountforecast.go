package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuheadcountforecastMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuheadcountforecastDud struct { 
    


    

}

// Buheadcountforecast
type Buheadcountforecast struct { 
    // Entities
    Entities []Buplanninggroupheadcountforecast `json:"entities"`


    // ReferenceStartDate - Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReferenceStartDate time.Time `json:"referenceStartDate"`

}

// String returns a JSON representation of the model
func (o *Buheadcountforecast) String() string {
     o.Entities = []Buplanninggroupheadcountforecast{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buheadcountforecast) MarshalJSON() ([]byte, error) {
    type Alias Buheadcountforecast

    if BuheadcountforecastMarshalled {
        return []byte("{}"), nil
    }
    BuheadcountforecastMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Buplanninggroupheadcountforecast `json:"entities"`
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        *Alias
    }{

        
        Entities: []Buplanninggroupheadcountforecast{{}},
        


        

        Alias: (*Alias)(u),
    })
}

