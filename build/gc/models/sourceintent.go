package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourceintentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourceintentDud struct { 
    


    


    


    


    

}

// Sourceintent
type Sourceintent struct { 
    // SourceIntentId - Id of the source intent
    SourceIntentId string `json:"sourceIntentId"`


    // SourceIntentName - Name of the source intent
    SourceIntentName string `json:"sourceIntentName"`


    // SourceType - Source type
    SourceType string `json:"sourceType"`


    // SourceId - Id of the source
    SourceId string `json:"sourceId"`


    // SourceName - Name of the source
    SourceName string `json:"sourceName"`

}

// String returns a JSON representation of the model
func (o *Sourceintent) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourceintent) MarshalJSON() ([]byte, error) {
    type Alias Sourceintent

    if SourceintentMarshalled {
        return []byte("{}"), nil
    }
    SourceintentMarshalled = true

    return json.Marshal(&struct {
        
        SourceIntentId string `json:"sourceIntentId"`
        
        SourceIntentName string `json:"sourceIntentName"`
        
        SourceType string `json:"sourceType"`
        
        SourceId string `json:"sourceId"`
        
        SourceName string `json:"sourceName"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

