package block

import (
	"net/http"
)

// our default block lists.
var blocklist = map[string]string{
	"StevenBlack": "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts",
	"MalwareDom":  "https://mirror1.malwaredomains.com/files/justdomains",
	"ZeusTracker": "https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist",
	"DisconAd":    "https://s3.amazonaws.com/lists.disconnect.me/simple_ad.txt",
	"HostsFile":   "https://hosts-file.net/ad_servers.txt",
}

func (b *Block) download() {
	//	client := &http.Client{Timeout: time.Second * 10}

	ok := 0
	for name, url := range blocklist {
		resp, err := http.Get(url)
		if err != nil {
			log.Warningf("Failed to download blocklist %q %q: %s", name, url, err)
			continue
		}
		if err := listRead(resp.Body, b.update); err != nil {
			log.Warningf("Failed to parse blocklist %q %q: %s", name, url, err)
		}
		ok += len(b.update)

		resp.Body.Close()
	}
	b.Lock()
	b.list = b.update
	b.update = make(map[string]struct{})
	b.Unlock()

	log.Infof("Block lists updates: %d domains added", ok)
}
