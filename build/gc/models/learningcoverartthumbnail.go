package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningcoverartthumbnailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningcoverartthumbnailDud struct { 
    


    

}

// Learningcoverartthumbnail
type Learningcoverartthumbnail struct { 
    // Resolution - Resolution of thumbnail
    Resolution string `json:"resolution"`


    // Url - The URL for the thumbnail
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Learningcoverartthumbnail) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningcoverartthumbnail) MarshalJSON() ([]byte, error) {
    type Alias Learningcoverartthumbnail

    if LearningcoverartthumbnailMarshalled {
        return []byte("{}"), nil
    }
    LearningcoverartthumbnailMarshalled = true

    return json.Marshal(&struct {
        
        Resolution string `json:"resolution"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

