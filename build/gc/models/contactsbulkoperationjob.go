package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactsbulkoperationjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactsbulkoperationjobDud struct { 
    Id string `json:"id"`


    State string `json:"state"`


    VarType string `json:"type"`


    TotalRecords int `json:"totalRecords"`


    CompletedRecords int `json:"completedRecords"`


    PercentComplete int `json:"percentComplete"`


    FailureReason Errorinfo `json:"failureReason"`


    DownloadURI string `json:"downloadURI"`


    SelfUri string `json:"selfUri"`

}

// Contactsbulkoperationjob
type Contactsbulkoperationjob struct { 
    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Contactsbulkoperationjob) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactsbulkoperationjob) MarshalJSON() ([]byte, error) {
    type Alias Contactsbulkoperationjob

    if ContactsbulkoperationjobMarshalled {
        return []byte("{}"), nil
    }
    ContactsbulkoperationjobMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

