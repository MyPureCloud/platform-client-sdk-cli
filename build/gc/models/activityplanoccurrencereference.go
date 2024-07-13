package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplanoccurrencereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplanoccurrencereferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Activityplanoccurrencereference
type Activityplanoccurrencereference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Activityplanoccurrencereference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplanoccurrencereference) MarshalJSON() ([]byte, error) {
    type Alias Activityplanoccurrencereference

    if ActivityplanoccurrencereferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivityplanoccurrencereferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

