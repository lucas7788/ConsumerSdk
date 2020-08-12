package base

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