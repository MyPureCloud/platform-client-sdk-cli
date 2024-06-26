package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DnclistdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DnclistdivisionviewDud struct { 
    Id string `json:"id"`


    


    


    ImportStatus Importstatus `json:"importStatus"`


    Size int `json:"size"`


    DncSourceType string `json:"dncSourceType"`


    


    SelfUri string `json:"selfUri"`

}

// Dnclistdivisionview
type Dnclistdivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    


    


    


    // ContactMethod - The contact method. Required if dncSourceType is rds.
    ContactMethod string `json:"contactMethod"`


    

}

// String returns a JSON representation of the model
func (o *Dnclistdivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dnclistdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Dnclistdivisionview

    if DnclistdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    DnclistdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        ContactMethod string `json:"contactMethod"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

