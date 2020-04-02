/* libzlintpq - run zlint from a PostgreSQL function
 * Written by Rob Stradling
 * Copyright (C) 2017-2020 Sectigo Limited
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
	"fmt"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v2"
	"github.com/zmap/zlint/v2/lint"
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
	registry := lint.GlobalRegistry()
	output := ""
	for k, v := range zlint_result.Results {
		switch v.Status {
			case lint.Notice: output += "N"
			case lint.Warn: output += "W"
			case lint.Error: output += "E"
			case lint.Fatal: output += "F"
			default: continue
		}
		output += ": " + registry.ByName(k).Description + "\n"
	}

	return output
}
