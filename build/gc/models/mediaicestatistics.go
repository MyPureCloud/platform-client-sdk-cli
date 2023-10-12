package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaicestatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaicestatisticsDud struct { 
    

}

// Mediaicestatistics
type Mediaicestatistics struct { 
    // SelectedPairs - The candidate pairs selected for the media stream
    SelectedPairs []Mediaiceselectedpair `json:"selectedPairs"`

}

// String returns a JSON representation of the model
func (o *Mediaicestatistics) String() string {
     o.SelectedPairs = []Mediaiceselectedpair{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaicestatistics) MarshalJSON() ([]byte, error) {
    type Alias Mediaicestatistics

    if MediaicestatisticsMarshalled {
        return []byte("{}"), nil
    }
    MediaicestatisticsMarshalled = true

    return json.Marshal(&struct {
        
        SelectedPairs []Mediaiceselectedpair `json:"selectedPairs"`
        *Alias
    }{

        
        SelectedPairs: []Mediaiceselectedpair{{}},
        

        Alias: (*Alias)(u),
    })
}

