package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneralprogramjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneralprogramjobrequestDud struct { 
    


    

}

// Generalprogramjobrequest
type Generalprogramjobrequest struct { 
    // Dialect - The dialect of the topics to link with the general program, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
    Dialect string `json:"dialect"`


    // Mode - The mode to use for the general program job, default value is Skip
    Mode string `json:"mode"`

}

// String returns a JSON representation of the model
func (o *Generalprogramjobrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generalprogramjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Generalprogramjobrequest

    if GeneralprogramjobrequestMarshalled {
        return []byte("{}"), nil
    }
    GeneralprogramjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Dialect string `json:"dialect"`
        
        Mode string `json:"mode"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

