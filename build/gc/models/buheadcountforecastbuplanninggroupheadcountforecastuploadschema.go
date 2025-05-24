package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuheadcountforecastbuplanninggroupheadcountforecastuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuheadcountforecastbuplanninggroupheadcountforecastuploadschemaDud struct { 
    


    

}

// Buheadcountforecastbuplanninggroupheadcountforecastuploadschema
type Buheadcountforecastbuplanninggroupheadcountforecastuploadschema struct { 
    // Entities
    Entities []Buplanninggroupheadcountforecastuploadschema `json:"entities"`


    // ReferenceStartDate - Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReferenceStartDate time.Time `json:"referenceStartDate"`

}

// String returns a JSON representation of the model
func (o *Buheadcountforecastbuplanninggroupheadcountforecastuploadschema) String() string {
     o.Entities = []Buplanninggroupheadcountforecastuploadschema{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buheadcountforecastbuplanninggroupheadcountforecastuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Buheadcountforecastbuplanninggroupheadcountforecastuploadschema

    if BuheadcountforecastbuplanninggroupheadcountforecastuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    BuheadcountforecastbuplanninggroupheadcountforecastuploadschemaMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Buplanninggroupheadcountforecastuploadschema `json:"entities"`
        
        ReferenceStartDate time.Time `json:"referenceStartDate"`
        *Alias
    }{

        
        Entities: []Buplanninggroupheadcountforecastuploadschema{{}},
        


        

        Alias: (*Alias)(u),
    })
}

