package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateuploadsourceurljobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateuploadsourceurljobresponseDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Createuploadsourceurljobresponse
type Createuploadsourceurljobresponse struct { 
    // Id - Id of the upload from URL job.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Createuploadsourceurljobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createuploadsourceurljobresponse) MarshalJSON() ([]byte, error) {
    type Alias Createuploadsourceurljobresponse

    if CreateuploadsourceurljobresponseMarshalled {
        return []byte("{}"), nil
    }
    CreateuploadsourceurljobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

