package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaendpointstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaendpointstatisticsDud struct { 
    


    


    


    


    

}

// Mediaendpointstatistics
type Mediaendpointstatistics struct { 
    // Trunk - Trunk information utilized when creating the media endpoint
    Trunk Mediastatisticstrunkinfo `json:"trunk"`


    // Station - Station information associated with media endpoint
    Station Namedentity `json:"station"`


    // User - User information associated media endpoint
    User Namedentity `json:"user"`


    // Ice - The ICE protocol statistics and details. Reference: https://www.rfc-editor.org/rfc/rfc5245
    Ice Mediaicestatistics `json:"ice"`


    // Rtp - Statistics of sent and received RTP. Reference: https://www.rfc-editor.org/rfc/rfc3550
    Rtp Mediartpstatistics `json:"rtp"`

}

// String returns a JSON representation of the model
func (o *Mediaendpointstatistics) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaendpointstatistics) MarshalJSON() ([]byte, error) {
    type Alias Mediaendpointstatistics

    if MediaendpointstatisticsMarshalled {
        return []byte("{}"), nil
    }
    MediaendpointstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        Trunk Mediastatisticstrunkinfo `json:"trunk"`
        
        Station Namedentity `json:"station"`
        
        User Namedentity `json:"user"`
        
        Ice Mediaicestatistics `json:"ice"`
        
        Rtp Mediartpstatistics `json:"rtp"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

