package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatetimeoffjobreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatetimeoffjobreferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Evaluatetimeoffjobreference
type Evaluatetimeoffjobreference struct { 
    


    // Status - The status of the job
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Evaluatetimeoffjobreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatetimeoffjobreference) MarshalJSON() ([]byte, error) {
    type Alias Evaluatetimeoffjobreference

    if EvaluatetimeoffjobreferenceMarshalled {
        return []byte("{}"), nil
    }
    EvaluatetimeoffjobreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

