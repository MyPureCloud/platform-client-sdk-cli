package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediastatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediastatisticsDud struct { 
    


    


    


    


    


    

}

// Mediastatistics
type Mediastatistics struct { 
    // CommunicationId
    CommunicationId string `json:"communicationId"`


    // DateStart - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // CreationMilliseconds - Relative milliseconds to create media session
    CreationMilliseconds int `json:"creationMilliseconds"`


    // PreferredRegion - Preferred media region
    PreferredRegion string `json:"preferredRegion"`


    // EffectiveRegion - Actual media region
    EffectiveRegion string `json:"effectiveRegion"`


    // MediaStatistics - Media statistics for each media endpoint
    MediaStatistics []Mediaendpointstatistics `json:"mediaStatistics"`

}

// String returns a JSON representation of the model
func (o *Mediastatistics) String() string {
    
    
    
    
    
     o.MediaStatistics = []Mediaendpointstatistics{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediastatistics) MarshalJSON() ([]byte, error) {
    type Alias Mediastatistics

    if MediastatisticsMarshalled {
        return []byte("{}"), nil
    }
    MediastatisticsMarshalled = true

    return json.Marshal(&struct {
        
        CommunicationId string `json:"communicationId"`
        
        DateStart time.Time `json:"dateStart"`
        
        CreationMilliseconds int `json:"creationMilliseconds"`
        
        PreferredRegion string `json:"preferredRegion"`
        
        EffectiveRegion string `json:"effectiveRegion"`
        
        MediaStatistics []Mediaendpointstatistics `json:"mediaStatistics"`
        *Alias
    }{

        


        


        


        


        


        
        MediaStatistics: []Mediaendpointstatistics{{}},
        

        Alias: (*Alias)(u),
    })
}

