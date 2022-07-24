# float64计算测试

test/amountdecimal_float64_test.go

```go
var (
	float64Min float64 = 0
	float64Max float64 = 99.9999999999999999999999999999999
	float64N   int     = 10000
	decimalInt int     = 22
)

=== RUN   TestFloat64
----------60.466028797961954 + 65.4829167211246 ------------------------
AmountDecimal result:<nil>, 125.9489455190865498934727 
AmountDecimal result2:<nil>, 125.9489455190865498934727 
Decimal result:<nil>, 125.9489455190865540000000 
--- FAIL: TestFloat64 (0.00s)
panic: AmountDecimal与Decimal不相等


=== RUN   TestFloat64TimeConsuming
开始时间： 1658659902071430100
结束时间： 1658659902099431600
耗时： 28001500

开始时间： 1658659902099431600
结束时间： 1658659902118596100
耗时： 19164500
--- PASS: TestFloat64TimeConsuming (0.05s)
PASS
```



```go
var (
	float64Min float64 = 0
	float64Max float64 = 99.9999999999999999999999999999999
	float64N   int     = 10000
	decimalInt int     = 15
)
----------60.466028797961954 + 65.4829167211246 ------------------------
AmountDecimal result:<nil>, 125.948945519086550 // 125.94894551908654|98934727 精度15，四舍五入进了一位
AmountDecimal result2:<nil>, 125.948945519086550 
Decimal result:<nil>, 125.948945519086554
--- FAIL: TestFloat64 (0.00s)
panic: AmountDecimal与Decimal不相等
```

