package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingprocessingerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingprocessingerrorDud struct { 
    InternalErrorCode string `json:"internalErrorCode"`


    Description string `json:"description"`

}

// Schedulingprocessingerror
type Schedulingprocessingerror struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Schedulingprocessingerror) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingprocessingerror) MarshalJSON() ([]byte, error) {
    type Alias Schedulingprocessingerror

    if SchedulingprocessingerrorMarshalled {
        return []byte("{}"), nil
    }
    SchedulingprocessingerrorMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

