package block

import (
	"bufio"
	"io"
	"strings"

	"github.com/miekg/dns"
)

// listRead parses two types of lists: a single and double column (host file like). We only care about the domain
// names. For the double column ones we only keep the second one.
func listRead(r io.Reader, list map[string]struct{}) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix("#", txt) {
			continue
		}
		var domain string
		flds := strings.Fields(scanner.Text())
		switch len(flds) {
		case 1:
			domain = dns.Fqdn(flds[0])
		case 2:
			domain = dns.Fqdn(flds[1])
		}
		// we only allow domains with more thna 2 dots, i.e. don't accidently block an entire TLD.
		if strings.Count(domain, ".") <= 2 {
			continue
		}
		list[domain] = struct{}{}
	}

	return scanner.Err()
}
