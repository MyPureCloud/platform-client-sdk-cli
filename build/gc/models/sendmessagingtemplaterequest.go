package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SendmessagingtemplaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SendmessagingtemplaterequestDud struct { 
    


    


    


    


    

}

// Sendmessagingtemplaterequest
type Sendmessagingtemplaterequest struct { 
    // ResponseId - A Response Management response identifier for a messaging template defined response
    ResponseId string `json:"responseId"`


    // Parameters - A list of Response Management response substitutions for the response's messaging template. (Deprecated) use bodyParameters instead.
    Parameters []Templateparameter `json:"parameters"`


    // HeaderParameters - A list of Response Management header parameter substitutions for the response's messaging template
    HeaderParameters []Templateparameter `json:"headerParameters"`


    // BodyParameters - A list of Response Management body parameter substitutions for the response's messaging template
    BodyParameters []Templateparameter `json:"bodyParameters"`


    // ButtonUrlParameters - A list of Response Management button url parameter substitutions for the response's messaging template
    ButtonUrlParameters []Templateparameter `json:"buttonUrlParameters"`

}

// String returns a JSON representation of the model
func (o *Sendmessagingtemplaterequest) String() string {
    
     o.Parameters = []Templateparameter{{}} 
     o.HeaderParameters = []Templateparameter{{}} 
     o.BodyParameters = []Templateparameter{{}} 
     o.ButtonUrlParameters = []Templateparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sendmessagingtemplaterequest) MarshalJSON() ([]byte, error) {
    type Alias Sendmessagingtemplaterequest

    if SendmessagingtemplaterequestMarshalled {
        return []byte("{}"), nil
    }
    SendmessagingtemplaterequestMarshalled = true

    return json.Marshal(&struct {
        
        ResponseId string `json:"responseId"`
        
        Parameters []Templateparameter `json:"parameters"`
        
        HeaderParameters []Templateparameter `json:"headerParameters"`
        
        BodyParameters []Templateparameter `json:"bodyParameters"`
        
        ButtonUrlParameters []Templateparameter `json:"buttonUrlParameters"`
        *Alias
    }{

        


        
        Parameters: []Templateparameter{{}},
        


        
        HeaderParameters: []Templateparameter{{}},
        


        
        BodyParameters: []Templateparameter{{}},
        


        
        ButtonUrlParameters: []Templateparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

