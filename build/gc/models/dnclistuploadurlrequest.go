package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DnclistuploadurlrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DnclistuploadurlrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Dnclistuploadurlrequest
type Dnclistuploadurlrequest struct { 
    // SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
    SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`


    // ContentType - The content type of the file to upload. Allows all types are text/csv, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
    ContentType string `json:"contentType"`


    // Id - Id of your dnc list to upload to
    Id string `json:"id"`


    // PhoneColumns - The column names from your file from which to upload dnc phone numbers.
    PhoneColumns string `json:"phoneColumns"`


    // EmailColumns - The column names from your file from which to upload dnc emails.
    EmailColumns string `json:"emailColumns"`


    // CustomExclusionColumns - The column names from your file from which to upload dnc custom exclusion column entries.
    CustomExclusionColumns string `json:"customExclusionColumns"`


    // ExpirationDateTimeColumn - The column name from your file to use as the dnc expiration date time.
    ExpirationDateTimeColumn string `json:"expirationDateTimeColumn"`


    // WhatsAppColumns - The column names from your file from which to upload dnc whatsapp.
    WhatsAppColumns string `json:"whatsAppColumns"`

}

// String returns a JSON representation of the model
func (o *Dnclistuploadurlrequest) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dnclistuploadurlrequest) MarshalJSON() ([]byte, error) {
    type Alias Dnclistuploadurlrequest

    if DnclistuploadurlrequestMarshalled {
        return []byte("{}"), nil
    }
    DnclistuploadurlrequestMarshalled = true

    return json.Marshal(&struct {
        
        SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`
        
        ContentType string `json:"contentType"`
        
        Id string `json:"id"`
        
        PhoneColumns string `json:"phoneColumns"`
        
        EmailColumns string `json:"emailColumns"`
        
        CustomExclusionColumns string `json:"customExclusionColumns"`
        
        ExpirationDateTimeColumn string `json:"expirationDateTimeColumn"`
        
        WhatsAppColumns string `json:"whatsAppColumns"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

