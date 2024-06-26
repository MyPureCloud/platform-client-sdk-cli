package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DnclistcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DnclistcreateDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    ImportStatus Importstatus `json:"importStatus"`


    Size int `json:"size"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dnclistcreate
type Dnclistcreate struct { 
    


    // Name - The name of the DncList.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    


    


    // DncSourceType - The type of the DncList.
    DncSourceType string `json:"dncSourceType"`


    // ContactMethod - The contact method. Required if dncSourceType is rds.
    ContactMethod string `json:"contactMethod"`


    // LoginId - A dnc.com loginId. Required if the dncSourceType is dnc.com.
    LoginId string `json:"loginId"`


    // CampaignId - A dnc.com campaignId. Optional if the dncSourceType is dnc.com.
    CampaignId string `json:"campaignId"`


    // DncCodes - The list of dnc.com codes to be treated as DNC. Required if the dncSourceType is dnc.com.
    DncCodes []string `json:"dncCodes"`


    // LicenseId - A gryphon license number. Required if the dncSourceType is gryphon.
    LicenseId string `json:"licenseId"`


    // Division - The division this DncList belongs to.
    Division Domainentityref `json:"division"`


    // CustomExclusionColumn - The column to evaluate exclusion against. Required if the dncSourceType is rds_custom.
    CustomExclusionColumn string `json:"customExclusionColumn"`


    

}

// String returns a JSON representation of the model
func (o *Dnclistcreate) String() string {
    
    
    
    
    
    
     o.DncCodes = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dnclistcreate) MarshalJSON() ([]byte, error) {
    type Alias Dnclistcreate

    if DnclistcreateMarshalled {
        return []byte("{}"), nil
    }
    DnclistcreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        DncSourceType string `json:"dncSourceType"`
        
        ContactMethod string `json:"contactMethod"`
        
        LoginId string `json:"loginId"`
        
        CampaignId string `json:"campaignId"`
        
        DncCodes []string `json:"dncCodes"`
        
        LicenseId string `json:"licenseId"`
        
        Division Domainentityref `json:"division"`
        
        CustomExclusionColumn string `json:"customExclusionColumn"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        DncCodes: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

