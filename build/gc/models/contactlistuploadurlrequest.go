package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistuploadurlrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistuploadurlrequestDud struct { 
    


    


    


    


    


    


    


    


    

}

// Contactlistuploadurlrequest
type Contactlistuploadurlrequest struct { 
    // SignedUrlTimeoutSeconds - The number of seconds the presigned URL is valid for (from 1 to 604800 seconds). If none provided, defaults to 600 seconds
    SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`


    // ContentType - The content type of the file to upload. Allows MIME types are text/csv, application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
    ContentType string `json:"contentType"`


    // Id - Id of your contact list to upload to
    Id string `json:"id"`


    // ContactIdName - The column name from your file to use as the contact id.
    ContactIdName string `json:"contactIdName"`


    // ImportTemplateId - Id of your import template to use.
    ImportTemplateId string `json:"importTemplateId"`


    // ListNamePrefix - String that will replace %N in the listNameFormat specified on the import template.
    ListNamePrefix string `json:"listNamePrefix"`


    // ClearSystemData - Whether to clear system data
    ClearSystemData bool `json:"clearSystemData"`


    // DivisionIdForTargetContactLists - Id of the division to be used for the creation of the target contact lists. If not provided, Home division will be used.
    DivisionIdForTargetContactLists string `json:"divisionIdForTargetContactLists"`


    // FileSpecificationTemplateId - File specification template ID
    FileSpecificationTemplateId string `json:"fileSpecificationTemplateId"`

}

// String returns a JSON representation of the model
func (o *Contactlistuploadurlrequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlistuploadurlrequest) MarshalJSON() ([]byte, error) {
    type Alias Contactlistuploadurlrequest

    if ContactlistuploadurlrequestMarshalled {
        return []byte("{}"), nil
    }
    ContactlistuploadurlrequestMarshalled = true

    return json.Marshal(&struct {
        
        SignedUrlTimeoutSeconds int `json:"signedUrlTimeoutSeconds"`
        
        ContentType string `json:"contentType"`
        
        Id string `json:"id"`
        
        ContactIdName string `json:"contactIdName"`
        
        ImportTemplateId string `json:"importTemplateId"`
        
        ListNamePrefix string `json:"listNamePrefix"`
        
        ClearSystemData bool `json:"clearSystemData"`
        
        DivisionIdForTargetContactLists string `json:"divisionIdForTargetContactLists"`
        
        FileSpecificationTemplateId string `json:"fileSpecificationTemplateId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

