package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingencryptionconfigurationlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingencryptionconfigurationlistingDud struct { 
    


    


    

}

// Recordingencryptionconfigurationlisting
type Recordingencryptionconfigurationlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Recordingencryptionconfiguration `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Recordingencryptionconfigurationlisting) String() string {
    
    
    
    
    
    
     o.Entities = []Recordingencryptionconfiguration{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingencryptionconfigurationlisting) MarshalJSON() ([]byte, error) {
    type Alias Recordingencryptionconfigurationlisting

    if RecordingencryptionconfigurationlistingMarshalled {
        return []byte("{}"), nil
    }
    RecordingencryptionconfigurationlistingMarshalled = true

    return json.Marshal(&struct { 
        Total int `json:"total"`
        
        Entities []Recordingencryptionconfiguration `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        
        *Alias
    }{
        

        

        

        
        Entities: []Recordingencryptionconfiguration{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

