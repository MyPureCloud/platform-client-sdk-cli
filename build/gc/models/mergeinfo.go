package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MergeinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MergeinfoDud struct { 
    Status string `json:"status"`


    VarError Mergeerror `json:"error"`


    DateMerged time.Time `json:"dateMerged"`

}

// Mergeinfo
type Mergeinfo struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Mergeinfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mergeinfo) MarshalJSON() ([]byte, error) {
    type Alias Mergeinfo

    if MergeinfoMarshalled {
        return []byte("{}"), nil
    }
    MergeinfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

