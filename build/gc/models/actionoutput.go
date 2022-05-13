package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionoutputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionoutputDud struct { 
    


    


    


    


    


    

}

// Actionoutput - Output definition of Action.
type Actionoutput struct { 
    // SuccessSchema - JSON schema that defines the transformed, successful result that will be sent back to the caller. If the 'flatten' query parameter is omitted or false, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
    SuccessSchema Jsonschemadocument `json:"successSchema"`


    // SuccessSchemaUri - URI to retrieve success schema
    SuccessSchemaUri string `json:"successSchemaUri"`


    // ErrorSchema - JSON schema that defines the body of response when request is not successful. If the 'flatten' query parameter is omitted or false, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
    ErrorSchema Jsonschemadocument `json:"errorSchema"`


    // ErrorSchemaUri - URI to retrieve error schema
    ErrorSchemaUri string `json:"errorSchemaUri"`


    // SuccessSchemaFlattened - JSON schema that defines the transformed, successful result that will be sent back to the caller. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
    SuccessSchemaFlattened Jsonschemadocument `json:"successSchemaFlattened"`


    // ErrorSchemaFlattened - JSON schema that defines the body of response when request is not successful. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
    ErrorSchemaFlattened interface{} `json:"errorSchemaFlattened"`

}

// String returns a JSON representation of the model
func (o *Actionoutput) String() string {
    
    
    
    
    
     o.ErrorSchemaFlattened = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionoutput) MarshalJSON() ([]byte, error) {
    type Alias Actionoutput

    if ActionoutputMarshalled {
        return []byte("{}"), nil
    }
    ActionoutputMarshalled = true

    return json.Marshal(&struct {
        
        SuccessSchema Jsonschemadocument `json:"successSchema"`
        
        SuccessSchemaUri string `json:"successSchemaUri"`
        
        ErrorSchema Jsonschemadocument `json:"errorSchema"`
        
        ErrorSchemaUri string `json:"errorSchemaUri"`
        
        SuccessSchemaFlattened Jsonschemadocument `json:"successSchemaFlattened"`
        
        ErrorSchemaFlattened interface{} `json:"errorSchemaFlattened"`
        *Alias
    }{

        


        


        


        


        


        
        ErrorSchemaFlattened: Interface{},
        

        Alias: (*Alias)(u),
    })
}

