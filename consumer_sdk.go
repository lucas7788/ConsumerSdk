package ConsumerSdk

import (
	"encoding/json"
	"github.com/kataras/go-errors"
	"github.com/ontio/ConsumerSdk/base"
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

func (this *ConsumerSdk) RequestAuthentication(ra base.RequestAuth) (err error) {
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
