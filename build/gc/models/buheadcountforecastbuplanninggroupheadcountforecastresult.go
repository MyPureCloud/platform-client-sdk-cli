package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuheadcountforecastbuplanninggroupheadcountforecastresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuheadcountforecastbuplanninggroupheadcountforecastresultDud struct { 
    


    

}

// Buheadcountforecastbuplanninggroupheadcountforecastresult
type Buheadcountforecastbuplanninggroupheadcountforecastresult struct { 
    // Entities
    Entities []Buplanninggroupheadcountforecastresult `json:"entities"`


    // ReferenceStartDate - Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReferenceStartDate time.Time `json:"referenceStartDate"`

}

// String returns a JSON representation of the model
func (o *Buheadcountforecastbuplanninggroupheadcountforecastresult) String() string {
     o.Entities = []Buplanninggroupheadcountforecastresult{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buheadcountforecastbuplanninggroupheadcountforecastresult) MarshalJSON() ([]byte, error) {
    type Alias Buheadcountforecastbuplanninggroupheadcountforecastresult

    if BuheadcountforecastbuplanninggroupheadcountforecastresultMarshalled {
        return []byte("{}"), nil
    }
    BuheadcountforecastbuplanninggroupheadcountforecastresultMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Buplanninggroupheadcountforecastresult `json:"entities"`
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        *Alias
    }{

        
        Entities: []Buplanninggroupheadcountforecastresult{{}},
        


        

        Alias: (*Alias)(u),
    })
}

