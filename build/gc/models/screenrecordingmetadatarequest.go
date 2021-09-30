package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenrecordingmetadatarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenrecordingmetadatarequestDud struct { 
    


    


    

}

// Screenrecordingmetadatarequest
type Screenrecordingmetadatarequest struct { 
    // ParticipantJid
    ParticipantJid string `json:"participantJid"`


    // RoomId
    RoomId string `json:"roomId"`


    // MetaData
    MetaData []Screenrecordingmetadata `json:"metaData"`

}

// String returns a JSON representation of the model
func (o *Screenrecordingmetadatarequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.MetaData = []Screenrecordingmetadata{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenrecordingmetadatarequest) MarshalJSON() ([]byte, error) {
    type Alias Screenrecordingmetadatarequest

    if ScreenrecordingmetadatarequestMarshalled {
        return []byte("{}"), nil
    }
    ScreenrecordingmetadatarequestMarshalled = true

    return json.Marshal(&struct { 
        ParticipantJid string `json:"participantJid"`
        
        RoomId string `json:"roomId"`
        
        MetaData []Screenrecordingmetadata `json:"metaData"`
        
        *Alias
    }{
        

        

        

        

        

        
        MetaData: []Screenrecordingmetadata{{}},
        

        
        Alias: (*Alias)(u),
    })
}

