package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatedhistoricalexportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatedhistoricalexportjobresponseDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Aggregatedhistoricalexportjobresponse
type Aggregatedhistoricalexportjobresponse struct { 
    // Id - The ID of the export job
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Aggregatedhistoricalexportjobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatedhistoricalexportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Aggregatedhistoricalexportjobresponse

    if AggregatedhistoricalexportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    AggregatedhistoricalexportjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

