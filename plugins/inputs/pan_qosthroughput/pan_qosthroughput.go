package qosthroughput

import (
	"github.com/PuerkitoBio/goquery"
	"strings"

	"github.com/irekromaniuk/telegraf"
	"github.com/irekromaniuk/telegraf/plugins/inputs"
	"strconv"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"regexp"
	"fmt"
)

type GetHTML func(url string) (string, error)

type Firewall struct {
	//  firewall's API key
	API string

	// IP address of firewall
	IP string

	// Names of interfaces and qos node-ids
	AE map[string]int

	// API call result
	HTML GetHTML
}

func (_ *Firewall) Description() string {
	return "Get throughput on firewall's interfaces"
}

const sampleConfig = `
  #
  ## firewall's API key
  api = "" # required
  ## IP address of firewall
  ip = "" # required
  ## Names of interfaces and node-ids
  int = {"ae1":1,"ae2":0,"ae3":0,}
`

func (_ *Firewall) SampleConfig() string {
	return sampleConfig
}
//SELECT "qos_throughput" FROM "qos_throughput" GROUP BY "class", "int" limit 3
func (p *Firewall) Gather(acc telegraf.Accumulator) error {
        var (
		tags string
		fields string
	)
	for k, v := range p.AE {
		out, err := p.HTML(p.IP + "&cmd=<show><qos><throughput>" + strconv.Itoa(v) + "</throughput><interface>" + k + "</interface></qos></show>" + p.API)
		if err != nil { return err }
		class := parseThroughput("result", out)
		for i, c := range class {
			//s, _ := strconv.Atoi(c)
			fmt.Println(c, i)
			tags = map[string]string{"class": c, "int": k}
			fields = map[string]interface{}{
				"qos_throughput": c,
			}
		}

		acc.AddFields("qos_throughput", fields, tags)
	}

	return nil
}

func getHTML (url string ) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil { return "", err }
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil { return "", err }
	resp.Body.Close()
	return string(htmlData), nil
}

func parseThroughput (tag string, htmlData string) []string {
	r := regexp.MustCompile("[^\\s]+")
	htmlCode := strings.NewReader(htmlData)
	doc, err := goquery.NewDocumentFromReader(htmlCode)
	if err != nil { log.Fatal(err) }
	//s := strings.Split(doc.Find(tag).Text()," ")
	s := r.FindAllString(doc.Find(tag).Text(),-1)
	class := []string{}
	for i:=2;i<=30;i+=4 {
		class = append(class,s[i])
	}
	return class
}

func init() {
	inputs.Add("pan_qosthroughput", func() telegraf.Input {
		return &Firewall{
			HTML: getHTML,
			API: "",
			IP: "",
			AE: {"ae1":1,"ae2":0,"ae3":0,},
		}
	})
}
