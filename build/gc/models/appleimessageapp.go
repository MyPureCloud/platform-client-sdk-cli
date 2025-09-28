package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleimessageappMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleimessageappDud struct { 
    


    


    

}

// Appleimessageapp
type Appleimessageapp struct { 
    // ApplicationName - Application Name.
    ApplicationName string `json:"applicationName"`


    // ApplicationId - Application ID.
    ApplicationId string `json:"applicationId"`


    // BundleId - Bundle ID.
    BundleId string `json:"bundleId"`

}

// String returns a JSON representation of the model
func (o *Appleimessageapp) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleimessageapp) MarshalJSON() ([]byte, error) {
    type Alias Appleimessageapp

    if AppleimessageappMarshalled {
        return []byte("{}"), nil
    }
    AppleimessageappMarshalled = true

    return json.Marshal(&struct {
        
        ApplicationName string `json:"applicationName"`
        
        ApplicationId string `json:"applicationId"`
        
        BundleId string `json:"bundleId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

