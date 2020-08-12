package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go/src/crypto/hmac"
	"strconv"
	"time"
)

func GenerateHmac(appId string, appKey string, method string, requestUri string, obj string) (res string, err error) {
	now := time.Now().Unix()
	timeStamp := strconv.FormatInt(now, 10)
	obj = `{"front_doc":"string","country":"EN","claim_context":"claim:sfp_idcard_authentication","auth_id":"111111111111111111111111111111","ar_ontid":"did:ont:Aejfo7ZX5PVpenRj23yChnyH64nf8T1zbu","description":"string","last_name":"string","doc_type":"id_card","doc_id":"12345678","ta_ontid":"did:ont:APc8FBdGYdzDtWrFp8q2BSUFX2HAnBuBna","back_doc":"string","encrp_flag":true,"owner_ontid":"did:ont:Aejfo7ZX5PVpenRj23yChnyH64nf8T1zbu","first_name":"string"}`
	temp := md5.Sum([]byte(obj))
	nonce := "test"
	bodyMD5Base64Str := base64.StdEncoding.EncodeToString(temp[:])
	fmt.Println("body:", bodyMD5Base64Str)
	rawData := appId + method + requestUri + timeStamp + nonce + bodyMD5Base64Str
	mac := hmac.New(sha256.New, []byte(appKey))
	mac.Write([]byte(rawData))
	fmt.Println("rawData:", rawData)
	sig := mac.Sum(nil)
	sigStr := base64.StdEncoding.EncodeToString(sig)
	res = fmt.Sprintf("ont:%s:%s:%s:%s", appId, sigStr, nonce, timeStamp)
	return
}
