# wrk -t2 -c12 -d60s --script=applyCredit.lua --latency http://10.96.140.18:8083/applyCredit
# wrk -t1 -c1 -d1s --script=applyCredit.lua --latency http://localhost:8211/applyCredit
wrk -t2 -c4 -d60s --latency http://10.96.140.18:8080/hello/smart
