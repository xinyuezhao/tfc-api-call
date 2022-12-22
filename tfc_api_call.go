package main

import (
	"fmt"
	"net/http"

	"github.com/ciscoecosystem/mso-go-client/client"
)

// func queryAgentIdByName() {

// }

func main() {
	clientND := client.GetClient("https://172.31.187.83", "admin", client.Password("ins3965!"), client.Insecure(true), client.Platform("nd"))
	url := "https://172.31.187.83/sedgeapi/v1/cisco-terraform/credentialsmgr"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error while building request to query credentials")
	}
	clientND.Do(req)

	// clientND.GetViaURL("https://172.31.187.83/sedgeapi/v1/cisco-terraform/credentialsmgr/api/terraform.argo.cisco.com/v1/credentials/cindy")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: tr}
	// // query agents inside agentpool tfc-app-test
	// req, err := http.NewRequest(http.MethodGet, "https://app.terraform.io/api/v2/agent-pools/apool-8PRRnPHJGJfyQaSx/agents", nil)
	// if err != nil {
	// 	fmt.Println("err ")
	// }
	// // use user token to access terraform cloud API
	// req.Header.Set("Authorization", "Bearer ZCUWZISXNFtWIg.atlasv1.vty2xgI8e0zuvzwgM9INeLvus2WYZPz5uziE1YU0UB27RiIDNunkXjFYxjlm7fDZxMc")
	// resp, e := client.Do(req)
	// if e != nil {
	// 	fmt.Println("err ")
	// }
	// defer resp.Body.Close()
	// // b, err := httputil.DumpResponse(resp, true)
	// // if err != nil {
	// // 	fmt.Println("err ")
	// // }
	// // fmt.Println("response content" + string(b))

	// // parse resp.Body

	// var body struct {
	// 	Data []struct {
	// 		Id         string
	// 		Attributes struct {
	// 			Name    string
	// 			Status  string
	// 			Version string
	// 		}
	// 	}
	// }

	// fmt.Println("response status " + resp.Status)

	// err = json.NewDecoder(resp.Body).Decode(&body)
	// if err != nil {
	// 	fmt.Println("err ")
	// }
	// // fmt.Println(body.Data[0].Id)

	// for _, agent := range body.Data {
	// 	fmt.Println("agent token " + agent.Id)
	// 	fmt.Println("agent name " + agent.Attributes.Name)
	// 	fmt.Println("agent status " + agent.Attributes.Status)
	// 	fmt.Println("agent version " + agent.Attributes.Version)
	// }

	// var data []map[string]interface{}

	// err = json.NewDecoder(j["data"]).Decode(&data)

	// agentId := "agent-JPsbet7zNVQ763SB"
	// url := fmt.Sprintf("https://app.terraform.io/api/v2/agents/%s", agentId)
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // use user token to access terraform cloud API
	// req.Header.Set("Authorization", "Bearer ZCUWZISXNFtWIg.atlasv1.vty2xgI8e0zuvzwgM9INeLvus2WYZPz5uziE1YU0UB27RiIDNunkXjFYxjlm7fDZxMc")
	// resp, e := client.Do(req)
	// if e != nil {
	// 	fmt.Println(e)
	// }
	// defer resp.Body.Close()
	// b, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(b))

	// type Components struct {
	// 	Components struct {
	// 		Terraform struct {
	// 			Credentials string
	// 		}
	// 	}
	// }
	// name := "terraform"
	// token := "testCindy"
	// jsonStr := fmt.Sprintf(`"components": {"terraform": {"sharedWith": ["cisco-terraform"],"credentials": {"token": "ai1yMKOzv3Mptg.atlasv1.lOseEHJzlB49Vz0fXTlFUFRGGTuugiP3040sr1MGGOkHgRqzQ9FrpiUJzyTH1DzzFTM"}}}`)
	// var jsonData = []byte(`{"components": {"terraform": {"sharedWith": ["cisco-terraform"],"credentials": {"token": "ai1yMKOzv3Mptg.atlasv1.lOseEHJzlB49Vz0fXTlFUFRGGTuugiP3040sr1MGGOkHgRqzQ9FrpiUJzyTH1DzzFTM"}}}}`)
	// req, err := http.NewRequest(http.MethodPost, "https://10.23.248.67/api/config/addcredentials", bytes.NewBuffer(jsonData))

	// credential := map[string]string{"token": token}
	// sharedWith := []string{"cisco-terraform"}
	// param := map[string]interface{}{
	// 	"credentials": map[string]string{"token": token},
	// 	"sharedWith":  []string{"cisco-terraform"},
	// }
	// // param := make(map[string]interface{})
	// // param["credentials"] = credential
	// // param["sharedWith"] = sharedWith
	// component := map[string]interface{}{
	// 	name: map[string]interface{}{
	// 		"credentials": map[string]string{"token": token},
	// 		"sharedWith":  []string{"cisco-terraform"},
	// 	},
	// }
	// payload := map[string]interface{}{
	// 	"components": component,
	// }
	// payload := map[string]interface{}{
	// 	"components": map[string]interface{}{
	// 		name: map[string]interface{}{
	// 			"credentials": map[string]string{"token": token},
	// 			"sharedWith":  []string{"cisco-terraform"},
	// 		},
	// 	},
	// }
	// payloadBuf := new(bytes.Buffer)
	// json.NewEncoder(payloadBuf).Encode(payload)
	// fmt.Println(payload)
	// req, err := http.NewRequest(http.MethodPost, "https://172.31.187.83/api/config/addcredentials", payloadBuf)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "AuthCookie=eyJhbGciOiJSUzI1NiIsImtpZCI6InJ5djd1bTRkamZvM3lqejAwb3h2aTQ5djNpcWMwZGNzIiwidHlwIjoiSldUIn0.eyJhdnBhaXIiOiJzaGVsbDpkb21haW5zPWFsbC9hZG1pbi8iLCJjbHVzdGVyIjoiNzY0ZTQ0MmQtMzgzMy0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiZXhwIjoxNjQyNTMwMDQwLCJpYXQiOjE2NDI1Mjg4NDAsImlkIjoiNDhkMTA1YmRmYmM0OWE1ZmNmMzlhMTBiOTYxMzg2ZTYxZGZlNDAwODVjYjAzMTVkODE4Yjc2MWM1NzM1ZGFmYSIsImlzcyI6Im5kIiwiaXNzLWhvc3QiOiIxNzIuMzEuMTg3LjgzIiwicmJhYyI6W3siZG9tYWluIjoiYWxsIiwicm9sZXMiOltbImFkbWluIiwiV3JpdGVQcml2Il0sWyJhcHAtdXNlciIsIlJlYWRQcml2Il1dLCJyb2xlc1IiOjE2Nzc3MjE2LCJyb2xlc1ciOjF9XSwic2Vzc2lvbmlkIjoicTlLcDI2VkZmM0RYT1NKWHh3WUI5bnNwIiwidXNlcmZsYWdzIjowLCJ1c2VyaWQiOjI1MDAyLCJ1c2VybmFtZSI6ImFkbWluIiwidXNlcnR5cGUiOiJsb2NhbCJ9.V_4ATToz2nPPPYbd5UFA4euhpUPP6F0RTJ9UX6-npxpAnEXebCKtsnNh-crt9Cu-_JyLtPm0n0xIwqskvLSXvWHKkci2tYhFaCOuLf6dwFwxOe6jZ6V-jtSJ7gHSRNqCtlm73jgYehSUkJw7s_c3ebO5HWuG2a2qWhf6k2TrXLS7sFR42JUGN7kZdElQpxFzNtAeoVWLt61DRGvA3EW9eyM8co2AQ8RnATp4IpbxKVJWZqRHYAl8llyNVm6cZUBa025k-i6QYZHLOHqJEIDQ1DK5XYMYuTO4soLsgvq-tyN-us4SLw60hV9w1Q4jDv76Wi0d7AvOkqcQVI3m6CZmyA")
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// b, err := httputil.DumpResponse(resp, true)
	// fmt.Printf(fmt.Sprintf("parsing response data %s", string(b)))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != 200 {
	// 	err := fmt.Errorf("error! Response content: %s", string(b))
	// 	fmt.Println(err)
	// }
	//query credentials
	// queryPayload := map[string]interface{}{
	// 	"components": map[string]interface{}{
	// 		name: map[string]string{},
	// 	},
	// }
	// queryPayloadBuf := new(bytes.Buffer)
	// json.NewEncoder(queryPayloadBuf).Encode(queryPayload)
	// queryReq, err := http.NewRequest(http.MethodPost, "https://172.31.187.83/api/config/getcredentials", queryPayloadBuf)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// queryReq.Header.Set("Content-Type", "application/json")
	// queryReq.Header.Set("Cookie", "AuthCookie=eyJhbGciOiJSUzI1NiIsImtpZCI6InJ5djd1bTRkamZvM3lqejAwb3h2aTQ5djNpcWMwZGNzIiwidHlwIjoiSldUIn0.eyJhdnBhaXIiOiJzaGVsbDpkb21haW5zPWFsbC9hZG1pbi8iLCJjbHVzdGVyIjoiNzY0ZTQ0MmQtMzgzMy0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiZXhwIjoxNjQyNjczNDQzLCJpYXQiOjE2NDI2NzIyNDMsImlkIjoiNDhkMTA1YmRmYmM0OWE1ZmNmMzlhMTBiOTYxMzg2ZTYxZGZlNDAwODVjYjAzMTVkODE4Yjc2MWM1NzM1ZGFmYSIsImlzcyI6Im5kIiwiaXNzLWhvc3QiOiIxNzIuMzEuMTg3LjgzIiwicmJhYyI6W3siZG9tYWluIjoiYWxsIiwicm9sZXMiOltbImFkbWluIiwiV3JpdGVQcml2Il0sWyJhcHAtdXNlciIsIlJlYWRQcml2Il1dLCJyb2xlc1IiOjE2Nzc3MjE2LCJyb2xlc1ciOjF9XSwic2Vzc2lvbmlkIjoiUWJPRUhwQUZjb1BsR05ORGNZPU1XNThDIiwidXNlcmZsYWdzIjowLCJ1c2VyaWQiOjI1MDAyLCJ1c2VybmFtZSI6ImFkbWluIiwidXNlcnR5cGUiOiJsb2NhbCJ9.pUdfKgTHLRxh7L-UgmYXMxh9mqlBOUWm_ZCIHZve-CB4gBpT3hDJDRZhYjY-FCJTcXPDp5B6oL-u9DH3HdpggbWbjBMi5Cx6earlI-ENNrS3iy2a2GaiZ6zra6BFM67I86xF3xrd7m7cNUjwWlAQZuVigMRPVoG7LcA1XYKkJ15iIWkYeupNOjledOhT36CYZeX5Fu0KVcXnHj4zcFTv32MmxX72PPxp7FPgR361Lu8-QxKo5XWyQIxqgf_jst2UXVuDpYqZFjZByJJIhofW8sqmymtjPz87Hv5kelZg-p4E5bbSTWupGkJyU9ujWAMwbxgppegwa_9eQkjekHnwuA")
	// resp, err := client.Do(queryReq)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// resBody, _ := ioutil.ReadAll(resp.Body)
	// respStr := string(resBody)
	// resBytes := []byte(respStr)
	// var jsonRes map[string]interface{}
	// e := json.Unmarshal(resBytes, &jsonRes)
	// if e != nil {
	// 	fmt.Println("error while parsing response from GetCredentials")
	// }
	// response := jsonRes["response"].([]interface{})[0].(map[string]interface{})
	// credentials := response["components"].(map[string]interface{})[name].(map[string]interface{})
	// token := credentials["credentials"]
	// configured := false
	// tokenExist := false
	// if token != nil {
	// 	configured = true
	// 	tokenStr := token.(map[string]interface{})["token"].(string)
	// 	if tokenStr != "" {
	// 		tokenExist = true
	// 	}
	// 	fmt.Printf(fmt.Sprintf("token is %s", tokenStr))
	// }
	// fmt.Printf(fmt.Sprintf("configured %v, tokenExist %v", configured, tokenExist))
	// 	type Credential struct {
	// 		Response []struct {
	// 			Components struct {
	// 				Terraform struct {
	// 					Credentials map[string]interface{}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	credential := Credential{}
	// 	fmt.Printf("before parsing resp.Body")
	// 	err = json.NewDecoder(resp.Body).Decode(&credential)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("after parsing resp.Body")
	// 	// fmt.Println("credential:" + credential.Response[0].Components.Terraform.Credentials + ":as shown")
	// 	res := credential.Response[0].Components.Terraform.Credentials
	// 	fmt.Println("%+v", res)
	// 	if token, ok := res["token"]; ok {
	// 		fmt.Println("token exist " + token.(string))
	// 	} else {
	// 		fmt.Println("no token")
	// 	}
}
