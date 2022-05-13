package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchdownloadjobsubmissionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchdownloadjobsubmissionresultDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Batchdownloadjobsubmissionresult
type Batchdownloadjobsubmissionresult struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Batchdownloadjobsubmissionresult) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchdownloadjobsubmissionresult) MarshalJSON() ([]byte, error) {
    type Alias Batchdownloadjobsubmissionresult

    if BatchdownloadjobsubmissionresultMarshalled {
        return []byte("{}"), nil
    }
    BatchdownloadjobsubmissionresultMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

