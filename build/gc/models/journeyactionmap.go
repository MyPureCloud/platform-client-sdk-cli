package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyactionmapMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyactionmapDud struct { 
    


    

}

// Journeyactionmap
type Journeyactionmap struct { 
    // Id - The ID of the actionMap in the Journey System which triggered this action
    Id string `json:"id"`


    // Version - The version number of the actionMap in the Journey System at the time this action was triggered
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Journeyactionmap) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyactionmap) MarshalJSON() ([]byte, error) {
    type Alias Journeyactionmap

    if JourneyactionmapMarshalled {
        return []byte("{}"), nil
    }
    JourneyactionmapMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

