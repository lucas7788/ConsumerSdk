package utils

import (
	"testing"
	"fmt"
	"github.com/ontio/ConsumerSdk/base"
	"encoding/json"
	"go/src/crypto/md5"
	"go/src/encoding/base64"
)

func TestGenerateHmac2(t *testing.T) {
	r := md5.Sum([]byte("test"))
	res := base64.StdEncoding.EncodeToString(r[:])
	fmt.Println(res)
}

func TestGenerateHmac(t *testing.T) {

	raw := `{
  "ar_ontid": "did:ont:Aejfo7ZX5PVpenRj23yChnyH64nf8T1zbu",
  "auth_id": "111111111111111111111111111111",
  "back_doc": "string",
  "claim_context": "claim:sfp_idcard_authentication",
  "country": "EN",
  "description": "string",
  "doc_id": "12345678",
  "doc_type": "id_card",
  "encrp_flag": true,
  "front_doc": "string",
  "mobile": "string",
  "name": "string",
  "owner_ontid": "did:ont:Aejfo7ZX5PVpenRj23yChnyH64nf8T1zbu",
  "photo": "string",
  "ta_ontid": "did:ont:APc8FBdGYdzDtWrFp8q2BSUFX2HAnBuBna"
}`

	ra := base.RequestAuth{
		TaOntId:"did:ont:APc8FBdGYdzDtWrFp8q2BSUFX2HAnBuBna",
		ArOntid:"did:ont:Aejfo7ZX5PVpenRj23yChnyH64nf8T1zbu",
		OwnerOntid:"did:ont:Aejfo7ZX5PVpenRj23yChnyH64nf8T1zbu",
		ClaimContext:"claim:sfp_idcard_authentication",
		AuthId:"111111111111111111111111111111",
		Description:"string",
		EncrpFlag:true,
		DocId:"12345678",
		DocType:"id_card",
		Country:"EN",
		FirstName:"string",
		LastName:"string",
		FrontDoc:"string",
		BackDoc:"string",
	}

	data, err := json.Marshal(ra)
	if string(data) != raw {
		fmt.Println("data:",string(data))
	}

	res,err := GenerateHmac("Ws6pndfG","Smp0NWRo4ndDo5FEPy/eaPy6Mrp9zg==","POST","/v1/kyc-data",string(data))
	fmt.Println(res)
	fmt.Println(err)
}
