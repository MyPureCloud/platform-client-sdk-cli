package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingencryptionconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingencryptionconfigurationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Recordingencryptionconfiguration
type Recordingencryptionconfiguration struct { 
    


    // Url - When keyConfigurationType is LocalKeyManager, this should be the url for decryption and must specify the path to where GenesysCloud can requests decryption. When keyConfigurationType is KmsSymmetric, this should be the arn to the key alias for the master key
    Url string `json:"url"`


    // ApiId - The api id for Hawk Authentication. Null if keyConfigurationType is KmsSymmetric
    ApiId string `json:"apiId"`


    // ApiKey - The api shared symmetric key used for hawk authentication. Null if keyConfigurationType is KmsSymmetric
    ApiKey string `json:"apiKey"`


    // KeyConfigurationType - Type should be LocalKeyManager or KmsSymmetric when create or update Key configurations; 'Native' for disabling configuration.
    KeyConfigurationType string `json:"keyConfigurationType"`


    // LastError - The error message related to the configuration
    LastError Errorbody `json:"lastError"`


    

}

// String returns a JSON representation of the model
func (o *Recordingencryptionconfiguration) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingencryptionconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Recordingencryptionconfiguration

    if RecordingencryptionconfigurationMarshalled {
        return []byte("{}"), nil
    }
    RecordingencryptionconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        ApiId string `json:"apiId"`
        
        ApiKey string `json:"apiKey"`
        
        KeyConfigurationType string `json:"keyConfigurationType"`
        
        LastError Errorbody `json:"lastError"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

