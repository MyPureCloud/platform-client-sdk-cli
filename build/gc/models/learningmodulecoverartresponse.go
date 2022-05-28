package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmodulecoverartresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmodulecoverartresponseDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    


    

}

// Learningmodulecoverartresponse - Learning module cover art response
type Learningmodulecoverartresponse struct { 
    


    


    // Url - The URL for the cover art
    Url string `json:"url"`


    // Thumbnails - Thumbnails for the cover art
    Thumbnails []Learningcoverartthumbnail `json:"thumbnails"`

}

// String returns a JSON representation of the model
func (o *Learningmodulecoverartresponse) String() string {
    
     o.Thumbnails = []Learningcoverartthumbnail{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmodulecoverartresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmodulecoverartresponse

    if LearningmodulecoverartresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmodulecoverartresponseMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        Thumbnails []Learningcoverartthumbnail `json:"thumbnails"`
        *Alias
    }{

        


        


        


        
        Thumbnails: []Learningcoverartthumbnail{{}},
        

        Alias: (*Alias)(u),
    })
}

