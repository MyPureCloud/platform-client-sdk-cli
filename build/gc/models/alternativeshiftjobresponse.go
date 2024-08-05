package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Alternativeshiftjobresponse
type Alternativeshiftjobresponse struct { 
    


    // Status - The status of the alternative shift job
    Status string `json:"status"`


    // VarType - The type of job
    VarType string `json:"type"`


    // DownloadUrl - The URL where completed results are available, only set if status == 'Complete'
    DownloadUrl string `json:"downloadUrl"`


    // VarError - Any error information, only set if the status == 'Error'
    VarError Errorbody `json:"error"`


    // ViewOffersResults - Schema template for deserializing data returned from the downloadUrl. Use if type == 'ListOffers' or 'SearchOffers'
    ViewOffersResults Alternativeshiftoffersviewresponsetemplate `json:"viewOffersResults"`


    // ViewTradesResults - Schema template for deserializing data returned from the downloadUrl. Use if type == 'ListUserTrades' or 'SearchTrades'
    ViewTradesResults Alternativeshifttradesviewresponsetemplate `json:"viewTradesResults"`


    // BulkUpdateTradesResults - Schema template for deserializing data returned from the downloadUrl. Use if type == 'BulkUpdateTrades'
    BulkUpdateTradesResults Alternativeshiftbulkupdatetradesresponsetemplate `json:"bulkUpdateTradesResults"`


    

}

// String returns a JSON representation of the model
func (o *Alternativeshiftjobresponse) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftjobresponse

    if AlternativeshiftjobresponseMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        VarError Errorbody `json:"error"`
        
        ViewOffersResults Alternativeshiftoffersviewresponsetemplate `json:"viewOffersResults"`
        
        ViewTradesResults Alternativeshifttradesviewresponsetemplate `json:"viewTradesResults"`
        
        BulkUpdateTradesResults Alternativeshiftbulkupdatetradesresponsetemplate `json:"bulkUpdateTradesResults"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

