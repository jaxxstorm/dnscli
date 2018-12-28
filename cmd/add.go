// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a DNS record",
	Long:  `Add a DNS record to a DNS server`,
	Run: func(cmd *cobra.Command, args []string) {

		s := strings.Split(record, ".")

		domain := fmt.Sprintf("%s.%s.", s[1], s[2])

		if len(s) < 3 {
			log.Fatal("Record not valid, please specify a fully qualified name")
		}

		// build a message
		newRR := fmt.Sprintf("%s %d %s %s", record, ttl, recordType, address)

		rr, err := dns.NewRR(newRR)
		if err != nil {
			fmt.Println("Whoops")
		}

		rrs := make([]dns.RR, 1)
		rrs[0] = rr

		m := new(dns.Msg)
		m.SetUpdate(domain)
		m.Insert(rrs)

		c := new(dns.Client)
		c.SingleInflight = true

		resp, _, err := c.Exchange(m, "localhost:53")

		if err != nil {
			log.Infof("error in dns.Client.Exchange: %s", err)
		}
		if resp != nil && resp.Rcode != dns.RcodeSuccess {
			log.Infof("Bad dns.Client.Exchange response: %s", resp)
		}

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
