package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DnclistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DnclistDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    ImportStatus Importstatus `json:"importStatus"`


    Size int `json:"size"`


    DncSourceType string `json:"dncSourceType"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dnclist
type Dnclist struct { 
    


    // Name - The name of the DncList.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    


    


    


    // ContactMethod - The contact method. Required if dncSourceType is rds.
    ContactMethod string `json:"contactMethod"`


    // LoginId - A dnc.com loginId. Required if the dncSourceType is dnc.com.
    LoginId string `json:"loginId"`


    // DncCodes - The list of dnc.com codes to be treated as DNC. Required if the dncSourceType is dnc.com.
    DncCodes []string `json:"dncCodes"`


    // LicenseId - A gryphon license number. Required if the dncSourceType is gryphon.
    LicenseId string `json:"licenseId"`


    // Division - The division this DncList belongs to.
    Division Domainentityref `json:"division"`


    

}

// String returns a JSON representation of the model
func (o *Dnclist) String() string {
    
    
    
    
     o.DncCodes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dnclist) MarshalJSON() ([]byte, error) {
    type Alias Dnclist

    if DnclistMarshalled {
        return []byte("{}"), nil
    }
    DnclistMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        ContactMethod string `json:"contactMethod"`
        
        LoginId string `json:"loginId"`
        
        DncCodes []string `json:"dncCodes"`
        
        LicenseId string `json:"licenseId"`
        
        Division Domainentityref `json:"division"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        
        DncCodes: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

