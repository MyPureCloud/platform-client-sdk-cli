package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenerateguidecontentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenerateguidecontentrequestDud struct { 
    


    

}

// Generateguidecontentrequest - Request body for generating the content of a guide
type Generateguidecontentrequest struct { 
    // Description - The description that you wish to use to generate the guide content from.
    Description string `json:"description"`


    // Url - The URL of the file you wish to use to generate the guide content from.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Generateguidecontentrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generateguidecontentrequest) MarshalJSON() ([]byte, error) {
    type Alias Generateguidecontentrequest

    if GenerateguidecontentrequestMarshalled {
        return []byte("{}"), nil
    }
    GenerateguidecontentrequestMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

