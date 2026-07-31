package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponsedivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponsedivisionviewDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Responsedivisionview - Division view of a response management response.
type Responsedivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // ResponseType - The response type represented by the response.
    ResponseType string `json:"responseType"`


    // Libraries - One or more libraries response is associated with.
    Libraries []Librarydivisionview `json:"libraries"`


    // Substitutions - Details about any text substitutions used in the texts for this response.
    Substitutions []Responsesubstitution `json:"substitutions"`


    // SubstitutionsSchema - Metadata about the text substitutions in json schema format.
    SubstitutionsSchema Jsonschemadocument `json:"substitutionsSchema"`


    // MessagingTemplate - An optional messaging template definition for responseType.MessagingTemplate.
    MessagingTemplate Messagingtemplate `json:"messagingTemplate"`


    // Form - Form template definition for responseType.Form.
    Form Form `json:"form"`


    

}

// String returns a JSON representation of the model
func (o *Responsedivisionview) String() string {
    
    
     o.Libraries = []Librarydivisionview{{}} 
     o.Substitutions = []Responsesubstitution{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responsedivisionview) MarshalJSON() ([]byte, error) {
    type Alias Responsedivisionview

    if ResponsedivisionviewMarshalled {
        return []byte("{}"), nil
    }
    ResponsedivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ResponseType string `json:"responseType"`
        
        Libraries []Librarydivisionview `json:"libraries"`
        
        Substitutions []Responsesubstitution `json:"substitutions"`
        
        SubstitutionsSchema Jsonschemadocument `json:"substitutionsSchema"`
        
        MessagingTemplate Messagingtemplate `json:"messagingTemplate"`
        
        Form Form `json:"form"`
        *Alias
    }{

        


        


        


        
        Libraries: []Librarydivisionview{{}},
        


        
        Substitutions: []Responsesubstitution{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

