package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherencebulkjobreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherencebulkjobreferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Wfmhistoricaladherencebulkjobreference
type Wfmhistoricaladherencebulkjobreference struct { 
    // Id - The ID of the historical adherence bulk job to listen for via notification or query using the jobs route
    Id string `json:"id"`


    // Status - The status of the historical adherence bulk job
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkjobreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherencebulkjobreference) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherencebulkjobreference

    if WfmhistoricaladherencebulkjobreferenceMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherencebulkjobreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

