package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyappMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyappDud struct { 
    


    


    


    

}

// Journeyapp
type Journeyapp struct { 
    // Name - Name of the application (e.g. mybankingapp).
    Name string `json:"name"`


    // Namespace - Namespace of the application (e.g. com.genesys.bancodinero).
    Namespace string `json:"namespace"`


    // Version - Version of the application (e.g. 5.9.27).
    Version string `json:"version"`


    // BuildNumber - Build number of the application (e.g. 701).
    BuildNumber string `json:"buildNumber"`

}

// String returns a JSON representation of the model
func (o *Journeyapp) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyapp) MarshalJSON() ([]byte, error) {
    type Alias Journeyapp

    if JourneyappMarshalled {
        return []byte("{}"), nil
    }
    JourneyappMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Namespace string `json:"namespace"`
        
        Version string `json:"version"`
        
        BuildNumber string `json:"buildNumber"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

