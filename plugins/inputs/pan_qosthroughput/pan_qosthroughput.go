package pan_qosthroughput

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	. "github.com/IrekRomaniuk/lib-go/slice2map"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
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
	INT []string

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
  int = ["ae1:1","ae2:0","ae3:0"]
`

func (_ *Firewall) SampleConfig() string {
	return sampleConfig
}
//SELECT "qos_throughput" FROM "qos_throughput" GROUP BY "class", "int" limit 1
func (p *Firewall) Gather(acc telegraf.Accumulator) error {
        var (
		tags map[string]string
		fields map[string]interface{}
	)
	intMap := Slice2Map(p.INT)
	addr := "https://" + p.IP + "/esp/restapi.esp?type=op"
	key := "&key=" + p.API
	for k, v := range intMap {  //http://stackoverflow.com/questions/38579485/golang-convert-slices-into-map
		out, err := p.HTML(addr + "&cmd=<show><qos><throughput>" + strconv.Itoa(v) + "</throughput><interface>" + k + "</interface></qos></show>" + key)
		if err != nil { return err }
		class, err := parseThroughput("result", out)
		if err != nil { return err }
		for i, c := range class {
			// Print class, throughput and interface name .i.e. 130784 3 ae1
			// fmt.Println(k, strconv.Itoa(i), c, k)
			tags = map[string]string{"class": strconv.Itoa(i), "int": k,}
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
	fmt.Printf("The url is %v", url)
	resp, err := client.Get(url)
	if err != nil { return "", err }
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil { return "", err }
	resp.Body.Close()
	return string(htmlData), nil
}

func parseThroughput (tag string, htmlData string) ([]string, error) {
	r := regexp.MustCompile("[^\\s]+")
	htmlCode := strings.NewReader(htmlData)
	doc, err := goquery.NewDocumentFromReader(htmlCode)
	if err != nil { return nil, err  }
	//s := strings.Split(doc.Find(tag).Text()," ")
	s := r.FindAllString(doc.Find(tag).Text(),-1)
	class := []string{}
	for i:=2;i<=30;i+=4 {
		class = append(class,s[i])
	}
	return class, nil
}

func init() {
	inputs.Add("pan_qosthroughput", func() telegraf.Input {
		return &Firewall{
			HTML: getHTML,
			API: "",
			IP: "",
			INT: []string{"ae1:1","ae2:0","ae3:0",},
		}
	})
}
