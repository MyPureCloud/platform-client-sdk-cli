package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediastatisticslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediastatisticslistingDud struct { 
    

}

// Mediastatisticslisting
type Mediastatisticslisting struct { 
    // Entities
    Entities []Mediastatistics `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Mediastatisticslisting) String() string {
     o.Entities = []Mediastatistics{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediastatisticslisting) MarshalJSON() ([]byte, error) {
    type Alias Mediastatisticslisting

    if MediastatisticslistingMarshalled {
        return []byte("{}"), nil
    }
    MediastatisticslistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Mediastatistics `json:"entities"`
        *Alias
    }{

        
        Entities: []Mediastatistics{{}},
        

        Alias: (*Alias)(u),
    })
}

