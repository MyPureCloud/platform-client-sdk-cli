package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionfilesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionfilesDud struct { 
    


    


    

}

// Sessionfiles
type Sessionfiles struct { 
    // MetaData - Metadata for the requested session
    MetaData Sessionmetadata `json:"metaData"`


    // Offered - Offered data for the requested session
    Offered Sessionmetricdata `json:"offered"`


    // AverageHandleTime - Average handle time data for the requested session
    AverageHandleTime Sessionmetricdata `json:"averageHandleTime"`

}

// String returns a JSON representation of the model
func (o *Sessionfiles) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionfiles) MarshalJSON() ([]byte, error) {
    type Alias Sessionfiles

    if SessionfilesMarshalled {
        return []byte("{}"), nil
    }
    SessionfilesMarshalled = true

    return json.Marshal(&struct {
        
        MetaData Sessionmetadata `json:"metaData"`
        
        Offered Sessionmetricdata `json:"offered"`
        
        AverageHandleTime Sessionmetricdata `json:"averageHandleTime"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

