https://www.grant.pizza/blog/test-build-modes/

Compile: `go test -c -o testbinary`

Debug exec: `dlv exec testbinary`

Break on unit test: `break switchers.TestSwitch`

Step through: `step` and `stepi`

Print registers: `regs`