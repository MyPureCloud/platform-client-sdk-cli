package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AggregatedsnapshotexportjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AggregatedsnapshotexportjobresponseDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Aggregatedsnapshotexportjobresponse
type Aggregatedsnapshotexportjobresponse struct { 
    // Id - The ID of the export job
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Aggregatedsnapshotexportjobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Aggregatedsnapshotexportjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Aggregatedsnapshotexportjobresponse

    if AggregatedsnapshotexportjobresponseMarshalled {
        return []byte("{}"), nil
    }
    AggregatedsnapshotexportjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

