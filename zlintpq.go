/* libzlintpq - run zlint from a PostgreSQL function
 * Written by Rob Stradling
 * Copyright (C) 2017 COMODO CA Limited
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint"
	"github.com/zmap/zlint/lints"
)

func Zlint_wrapper(b64_cert string) string {
	der_cert, err := base64.StdEncoding.DecodeString(b64_cert)
	if err != nil {
		return fmt.Sprintf("F: %s", err)
	}

	cert, err := x509.ParseCertificate(der_cert)
	if err != nil {
		return fmt.Sprintf("F: %s", err)
	}

	zlint_result := zlint.LintCertificate(cert)
	json_result, err := json.Marshal(zlint_result.ZLint)
	if err != nil {
		return fmt.Sprintf("F: %s", err)
	}

	var f interface{}
	err = json.Unmarshal(json_result, &f)
	if err != nil {
		return fmt.Sprintf("F: %s", err)
	}

	m := f.(map[string]interface{})
	output := ""
	for k, v := range m {
		switch vv := v.(type) {
		default:
			switch fmt.Sprintf("%v", vv.(map[string]interface{})["result"]) {
				case "info": output += "N"
				case "warn": output += "W"
				case "error": output += "E"
				case "fatal": output += "F"
				default: continue
			}
			output += ": " + lints.Lints[k].Description + "\n"
		}
	}

	return output
}
