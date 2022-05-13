package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagingtemplaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagingtemplaterequestDud struct { 
    


    

}

// Messagingtemplaterequest
type Messagingtemplaterequest struct { 
    // ResponseId - A Response Management response identifier for a messaging template defined response
    ResponseId string `json:"responseId"`


    // Parameters - A list of Response Management response substitutions for the response's messaging template
    Parameters []Templateparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Messagingtemplaterequest) String() string {
    
     o.Parameters = []Templateparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagingtemplaterequest) MarshalJSON() ([]byte, error) {
    type Alias Messagingtemplaterequest

    if MessagingtemplaterequestMarshalled {
        return []byte("{}"), nil
    }
    MessagingtemplaterequestMarshalled = true

    return json.Marshal(&struct {
        
        ResponseId string `json:"responseId"`
        
        Parameters []Templateparameter `json:"parameters"`
        *Alias
    }{

        


        
        Parameters: []Templateparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

