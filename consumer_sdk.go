package ConsumerSdk

import (
	"encoding/json"
	"github.com/kataras/go-errors"
	"github.com/ontio/ConsumerSdk/config"
	"github.com/ontio/ConsumerSdk/forward"
	"github.com/ontio/ontology-go-sdk"
)

type ConsumerSdk struct {
	ontPassAddr string
	ontSdk      *ontology_go_sdk.OntologySdk
}

func NewConsumerSdk(ontPassAddr string, ontSdk *ontology_go_sdk.OntologySdk) *ConsumerSdk {
	return &ConsumerSdk{
		ontPassAddr: ontPassAddr,
		ontSdk:      ontSdk,
	}
}

func (this *ConsumerSdk) QueryIssuerList() (issuerList []string, err error) {
	if this.ontPassAddr == "" {
		err = errors.New("ontPassAddr is not set")
		return
	}
	res, err := forward.Get(this.ontPassAddr + config.QueryIssuerList)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &issuerList)
	if err != nil {
		return
	}
	return
}

func (this *ConsumerSdk) QueryTemplateByIssuerId(issuerId string) (template map[string]interface{}, err error) {
	if this.ontPassAddr == "" {
		err = errors.New("ontPassAddr is not set")
		return
	}
	res, err := forward.Get(this.ontPassAddr + config.QueryIssuerTemplateByIssuerId)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &template)
	if err != nil {
		return
	}
	return
}

type RequestAuth struct {
	TaOntId      string `json:"ta_ontid"`      //ONTTA的ONT ID
	ArOntid      string `json:"ar_ontid"`      //认证请求方的ONT ID
	OwnerOntid   string `json:"owner_ontid"`   //用户的ONT ID
	ClaimContext string `json:"claim_context"` //可信声明模板标识
	AuthId       string `json:"auth_id"`       //认证需求方的认证编号
	Description  string `json:"description"`   //
	EncrpFlag    bool   `json:"encrp_flag"`    // true:claim加密传输  false:claim不加密传输
	DocId        string `json:"doc_id"`        //证件编号
	DocType      string `json:"doc_type"`      //证件类型。passport:护照 id_card：身份证 driving_license：驾照
	Country      string `json:"country"`       //国籍
	FirstName    string `json:"first_name"`    // 名字
	LastName     string `json:"last_name"`     // 名字
	FrontDoc     string `json:"front_doc"`     //正面
	BackDoc      string `json:"back_doc"`      //反面
}

func (this *ConsumerSdk) RequestAuthentication(ra RequestAuth) (err error) {
	if this.ontPassAddr == "" {
		err = errors.New("ontPassAddr is not set")
		return
	}
	data, err := json.Marshal(ra)
	if err != nil {
		return
	}
	_, err = forward.PostJSONRequest(this.ontPassAddr+config.RequestAuthentication, data, nil)
	if err != nil {
		return
	}
	return
}

func (this *ConsumerSdk) VerifyCredential(credential *ontology_go_sdk.VerifiableCredential) (err error) {
	err = this.ontSdk.Credential.VerifyExpirationDate(credential)
	if err != nil {
		return
	}
	err = this.ontSdk.Credential.VerifyIssuanceDate(credential)
	if err != nil {
		return
	}
	err = this.ontSdk.Credential.VerifyIssuerSignature(credential)
	if err != nil {
		return
	}
	err = this.ontSdk.Credential.VerifyStatus(credential)
	if err != nil {
		return
	}
	return
}
