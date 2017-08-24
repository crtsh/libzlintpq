package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/lints"
	"github.com/zmap/zlint/zlint"
)

func Zlint_wrapper(b64_cert string) string {
	der_cert, err := base64.StdEncoding.DecodeString(b64_cert)
	if err != nil {
		return "ERROR"
	}

	cert, err := x509.ParseCertificate(der_cert)
	if err != nil {
		return "ERROR"
	}

	zlint_result := zlint.ZLintResultTestHandler(cert)
	json_result, err := json.Marshal(zlint_result.ZLint)
	if err != nil {
		return "ERROR"
	}

	var f interface{}
	err = json.Unmarshal(json_result, &f)
	if err != nil {
		return "ERROR"
	}

	m := f.(map[string]interface{})
	output := ""
	for k, v := range m {
		switch vv := v.(type) {
		default:
			switch fmt.Sprintf("%v", vv.(map[string]interface{})["result"]) {
				case "4": output += "N"
				case "5": output += "W"
				case "6": output += "E"
				case "7": output += "F"
				default: continue
			}
			output += ": " + lints.Lints[k].Description + "\n"
		}
	}

	return output
}
