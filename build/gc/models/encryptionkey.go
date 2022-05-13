package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EncryptionkeyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EncryptionkeyDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Encryptionkey
type Encryptionkey struct { 
    


    // Name
    Name string `json:"name"`


    // CreateDate - create date of the key pair. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreateDate time.Time `json:"createDate"`


    // KeydataSummary - key data summary (base 64 encoded public key)
    KeydataSummary string `json:"keydataSummary"`


    // User - user that requested generation of public key
    User User `json:"user"`


    // LocalEncryptionConfiguration - Local configuration
    LocalEncryptionConfiguration Localencryptionconfiguration `json:"localEncryptionConfiguration"`


    // KeyConfigurationType - Key type used in this configuration
    KeyConfigurationType string `json:"keyConfigurationType"`


    // KmsKeyArn - ARN of internal key to be wrapped by AWS KMS Symmetric key
    KmsKeyArn string `json:"kmsKeyArn"`


    

}

// String returns a JSON representation of the model
func (o *Encryptionkey) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Encryptionkey) MarshalJSON() ([]byte, error) {
    type Alias Encryptionkey

    if EncryptionkeyMarshalled {
        return []byte("{}"), nil
    }
    EncryptionkeyMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        CreateDate time.Time `json:"createDate"`
        
        KeydataSummary string `json:"keydataSummary"`
        
        User User `json:"user"`
        
        LocalEncryptionConfiguration Localencryptionconfiguration `json:"localEncryptionConfiguration"`
        
        KeyConfigurationType string `json:"keyConfigurationType"`
        
        KmsKeyArn string `json:"kmsKeyArn"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

