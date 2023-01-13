package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkqueryDud struct { 
    


    

}

// Wfmhistoricaladherencebulkquery
type Wfmhistoricaladherencebulkquery struct { 
    // Items - The historical adherence items to query
    Items []Wfmhistoricaladherencebulkitem `json:"items"`


    // TimeZone - The time zone, in olson format, to use in defining days when computing adherence. The results will be returned as UTC timestamps regardless of the time zone input.
    TimeZone string `json:"timeZone"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkquery) String() string {
     o.Items = []Wfmhistoricaladherencebulkitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkquery) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkquery

    if WfmhistoricaladherencebulkqueryMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkqueryMarshalled = true

    return json.Marshal(&struct {
        
        Items []Wfmhistoricaladherencebulkitem `json:"items"`
        
        TimeZone string `json:"timeZone"`
        *Alias
    }{

        
        Items: []Wfmhistoricaladherencebulkitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

