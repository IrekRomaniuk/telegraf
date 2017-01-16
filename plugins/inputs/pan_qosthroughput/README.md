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
tags: class=0, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    0

name: qos_throughput
tags: class=0, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    0

name: qos_throughput
tags: class=0, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

name: qos_throughput
tags: class=1, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    0

name: qos_throughput
tags: class=1, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    0

name: qos_throughput
tags: class=1, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

name: qos_throughput
tags: class=2, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    0

name: qos_throughput
tags: class=2, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    0

name: qos_throughput
tags: class=2, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

name: qos_throughput
tags: class=3, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    210261

name: qos_throughput
tags: class=3, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    31886

name: qos_throughput
tags: class=3, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    19330

name: qos_throughput
tags: class=4, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    0

name: qos_throughput
tags: class=4, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    0

name: qos_throughput
tags: class=4, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

name: qos_throughput
tags: class=5, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    0

name: qos_throughput
tags: class=5, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    0

name: qos_throughput
tags: class=5, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

name: qos_throughput
tags: class=6, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    600

name: qos_throughput
tags: class=6, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    20416

name: qos_throughput
tags: class=6, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

name: qos_throughput
tags: class=7, int=ae1
time                    qos_throughput
----                    --------------
2017-01-16T14:36:02Z    11

name: qos_throughput
tags: class=7, int=ae2
time                    qos_throughput
----                    --------------
2017-01-16T14:36:03Z    1

name: qos_throughput
tags: class=7, int=ae3
time                    qos_throughput
----                    --------------
2017-01-16T14:36:05Z    0

> 
```
