package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserscheduleadherencelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserscheduleadherencelistingDud struct { 
    


    DownloadUrl string `json:"downloadUrl"`

}

// Userscheduleadherencelisting
type Userscheduleadherencelisting struct { 
    // Entities
    Entities []Userscheduleadherence `json:"entities"`


    

}

// String returns a JSON representation of the model
func (o *Userscheduleadherencelisting) String() string {
    
    
     o.Entities = []Userscheduleadherence{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userscheduleadherencelisting) MarshalJSON() ([]byte, error) {
    type Alias Userscheduleadherencelisting

    if UserscheduleadherencelistingMarshalled {
        return []byte("{}"), nil
    }
    UserscheduleadherencelistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Userscheduleadherence `json:"entities"`
        
        
        
        *Alias
    }{
        

        
        Entities: []Userscheduleadherence{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

