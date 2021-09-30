package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    CreatedBy User `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Response - Contains information about a response.
type Response struct { 
    


    // Name
    Name string `json:"name"`


    // Version - Version number required for updates.
    Version int `json:"version"`


    // Libraries - One or more libraries response is associated with.
    Libraries []Domainentityref `json:"libraries"`


    // Texts - One or more texts associated with the response.
    Texts []Responsetext `json:"texts"`


    


    


    // InteractionType - The interaction type for this response.
    InteractionType string `json:"interactionType"`


    // Substitutions - Details about any text substitutions used in the texts for this response.
    Substitutions []Responsesubstitution `json:"substitutions"`


    // SubstitutionsSchema - Metadata about the text substitutions in json schema format.
    SubstitutionsSchema Jsonschemadocument `json:"substitutionsSchema"`


    // ResponseType - The response type represented by the response.
    ResponseType string `json:"responseType"`


    // MessagingTemplate - An optional messaging template definition for responseType.MessagingTemplate.
    MessagingTemplate Messagingtemplate `json:"messagingTemplate"`


    // Assets - Assets used in the response
    Assets []Addressableentityref `json:"assets"`


    

}

// String returns a JSON representation of the model
func (o *Response) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Libraries = []Domainentityref{{}} 
    
    
    
     o.Texts = []Responsetext{{}} 
    
    
    
    
    
    
    
    
    
    
    
     o.Substitutions = []Responsesubstitution{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Assets = []Addressableentityref{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Response) MarshalJSON() ([]byte, error) {
    type Alias Response

    if ResponseMarshalled {
        return []byte("{}"), nil
    }
    ResponseMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Libraries []Domainentityref `json:"libraries"`
        
        Texts []Responsetext `json:"texts"`
        
        
        
        
        
        InteractionType string `json:"interactionType"`
        
        Substitutions []Responsesubstitution `json:"substitutions"`
        
        SubstitutionsSchema Jsonschemadocument `json:"substitutionsSchema"`
        
        ResponseType string `json:"responseType"`
        
        MessagingTemplate Messagingtemplate `json:"messagingTemplate"`
        
        Assets []Addressableentityref `json:"assets"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Libraries: []Domainentityref{{}},
        

        

        
        Texts: []Responsetext{{}},
        

        

        

        

        

        

        

        

        
        Substitutions: []Responsesubstitution{{}},
        

        

        

        

        

        

        

        

        
        Assets: []Addressableentityref{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

