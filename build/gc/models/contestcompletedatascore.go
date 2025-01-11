package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestcompletedatascoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestcompletedatascoreDud struct { 
    

}

// Contestcompletedatascore
type Contestcompletedatascore struct { 
    // Score - The Contest score
    Score float64 `json:"score"`

}

// String returns a JSON representation of the model
func (o *Contestcompletedatascore) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestcompletedatascore) MarshalJSON() ([]byte, error) {
    type Alias Contestcompletedatascore

    if ContestcompletedatascoreMarshalled {
        return []byte("{}"), nil
    }
    ContestcompletedatascoreMarshalled = true

    return json.Marshal(&struct {
        
        Score float64 `json:"score"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

