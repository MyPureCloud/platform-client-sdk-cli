package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaresultDud struct { 
    


    

}

// Mediaresult
type Mediaresult struct { 
    // MediaUri
    MediaUri string `json:"mediaUri"`


    // WaveformData
    WaveformData []float32 `json:"waveformData"`

}

// String returns a JSON representation of the model
func (o *Mediaresult) String() string {
    
    
    
    
    
    
     o.WaveformData = []float32{0.0} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaresult) MarshalJSON() ([]byte, error) {
    type Alias Mediaresult

    if MediaresultMarshalled {
        return []byte("{}"), nil
    }
    MediaresultMarshalled = true

    return json.Marshal(&struct { 
        MediaUri string `json:"mediaUri"`
        
        WaveformData []float32 `json:"waveformData"`
        
        *Alias
    }{
        

        

        

        
        WaveformData: []float32{0.0},
        

        
        Alias: (*Alias)(u),
    })
}

