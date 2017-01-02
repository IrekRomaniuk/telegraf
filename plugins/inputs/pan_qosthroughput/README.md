# PaloAlto firewall qos throughput monitor input plugin

This input plugin will measures the qos throughput

### Configuration:
```
      ## firewall's API key
      api = "" # required
      ## IP address of firewall
      ip = "" # required
      ## Names of interfaces and node-ids
      int = ["ae1:1","ae2:0","ae3:0",]	
```
### Measurements & Fields:
- qos_throughput (kbps, per qos class and interface). Tested on PAN Model 5060 s/w 7.1.2
```
admin@PA-5060(active)> show qos throughput 
  <value>  <0-65535> Show throughput (last 3 seconds) of all classes under given node-id

admin@PA-5060(active)> show qos throughput 1 interface 
  ae1      ae1
  ae2      ae2
  ae3      ae3
  <value>  Show for given interface

admin@PA-5060(active)> show qos throughput 1 interface ae1
Class 1              0 kbps
Class 2              0 kbps
Class 3              0 kbps
Class 4         502811 kbps
Class 5              0 kbps
Class 6              0 kbps
Class 7            190 kbps
Class 8             39 kbps
```
	
### Tags:
- class, int

### Example Output:
```
> SELECT "qos_throughput" FROM "qos_throughput" GROUP BY "class", "int" limit 1
name: qos_throughput
tags: class=3, int=ae2
time                    qos_throughput
----                    --------------
1482767284000000000     12426
```
