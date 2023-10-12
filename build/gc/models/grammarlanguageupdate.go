package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GrammarlanguageupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GrammarlanguageupdateDud struct { 
    


    

}

// Grammarlanguageupdate
type Grammarlanguageupdate struct { 
    // VoiceFileMetadata - Additional information about the associated voice file
    VoiceFileMetadata Grammarlanguagefilemetadata `json:"voiceFileMetadata"`


    // DtmfFileMetadata - Additional information about the associated dtmf file
    DtmfFileMetadata Grammarlanguagefilemetadata `json:"dtmfFileMetadata"`

}

// String returns a JSON representation of the model
func (o *Grammarlanguageupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Grammarlanguageupdate) MarshalJSON() ([]byte, error) {
    type Alias Grammarlanguageupdate

    if GrammarlanguageupdateMarshalled {
        return []byte("{}"), nil
    }
    GrammarlanguageupdateMarshalled = true

    return json.Marshal(&struct {
        
        VoiceFileMetadata Grammarlanguagefilemetadata `json:"voiceFileMetadata"`
        
        DtmfFileMetadata Grammarlanguagefilemetadata `json:"dtmfFileMetadata"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

