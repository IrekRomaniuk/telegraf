package pan_qosthroughput

import (

	"testing"

	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"fmt"
)

func mockGetHTML1(url string) (string, error) {
	return HTMLThroughput1, nil
}

// Test that Gather function works on HTML1
func TestGather(t *testing.T) {
	var acc testutil.Accumulator
	p := Firewall{
		API: "",
		IP: "",
		INT: []string{"ae1:1",},
		HTML: mockGetHTML1,
	}

	err := p.Gather(&acc)
	assert.NoError(t, err)
	metric, ok := acc.Get("qos_throughput")
	require.True(t, ok)
	//fmt.Println(len(metric.Tags))  // 2
	qos_throughput := metric.Fields["qos_throughput"]
	//fmt.Println(metric.Tags["class"],metric.Tags["int"]) //7 ae1
	fmt.Printf("qos_throughput is %s\n", qos_throughput)
	tags := map[string]string{"class": "7", "int": "ae1",}
	fields := map[string]interface{}{}
	fields["qos_throughput"] = qos_throughput
	/*fields := map[string]interface{}{
		"qos_throughput": "13",
	}*/
	/*assert.False(t, acc.HasMeasurement("qos_throughput"),
		"Missing qos_throughput measurement")*/
	acc.AssertContainsTaggedFields(t, "qos_throughput", fields, tags)
}

var HTMLThroughput1 = `<response status="success">
<result>
Class 1 0 kbps Class 2 0 kbps Class 3 0 kbps Class 4 130784 kbps Class 5 0 kbps Class 6 0 kbps Class 7 20 kbps Class 8 13 kbps
</result>
</response>`

var HTMLThroughput2 = `<response status="success"><result>Class 1              0 kbps
Class 2              0 kbps
Class 3              0 kbps
Class 4         130784 kbps
Class 5              0 kbps
Class 6              0 kbps
Class 7             20 kbps
Class 8             13 kbps
</result></response>`



