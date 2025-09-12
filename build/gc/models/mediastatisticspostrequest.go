package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediastatisticspostrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediastatisticspostrequestDud struct { 
    


    


    


    


    

}

// Mediastatisticspostrequest
type Mediastatisticspostrequest struct { 
    // SourceType - Source type of media endpoint
    SourceType string `json:"sourceType"`


    // ClientInfo - Client information associated with media endpoint
    ClientInfo Mediastatisticsclientinfo `json:"clientInfo"`


    // Rtp - Statistics of sent and received RTP. Reference: https://www.rfc-editor.org/rfc/rfc3550
    Rtp Mediartpstatistics `json:"rtp"`


    // ReconnectAttempts - Media reconnect attempt count
    ReconnectAttempts int `json:"reconnectAttempts"`


    // DateCreated - Media endpoint statistics creation time. Format: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`

}

// String returns a JSON representation of the model
func (o *Mediastatisticspostrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediastatisticspostrequest) MarshalJSON() ([]byte, error) {
    type Alias Mediastatisticspostrequest

    if MediastatisticspostrequestMarshalled {
        return []byte("{}"), nil
    }
    MediastatisticspostrequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceType string `json:"sourceType"`
        
        ClientInfo Mediastatisticsclientinfo `json:"clientInfo"`
        
        Rtp Mediartpstatistics `json:"rtp"`
        
        ReconnectAttempts int `json:"reconnectAttempts"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

