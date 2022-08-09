package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AvailablemediatypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AvailablemediatypeDud struct { 
    


    

}

// Availablemediatype
type Availablemediatype struct { 
    // MediaType - Name of the available media type
    MediaType string `json:"mediaType"`


    // AvailableSubTypes - List of available subtypes for this media type
    AvailableSubTypes []string `json:"availableSubTypes"`

}

// String returns a JSON representation of the model
func (o *Availablemediatype) String() string {
    
     o.AvailableSubTypes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Availablemediatype) MarshalJSON() ([]byte, error) {
    type Alias Availablemediatype

    if AvailablemediatypeMarshalled {
        return []byte("{}"), nil
    }
    AvailablemediatypeMarshalled = true

    return json.Marshal(&struct {
        
        MediaType string `json:"mediaType"`
        
        AvailableSubTypes []string `json:"availableSubTypes"`
        *Alias
    }{

        


        
        AvailableSubTypes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

