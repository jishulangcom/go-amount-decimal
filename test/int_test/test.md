





```go
var (
	intMin     int = math.MinInt
	intMax     int = math.MaxInt64
	intNub     int     = 10000
	intDecimal int     = 20
)


----------3646365244906996398 / 1268292630671260171 ------------------------
AmountDecimal  result:<nil>, 2.87501886924716320413 
AmountDecimal result2:<nil>, 2.87501886924716320413 
Decimal result:<nil>, 2.87501886924716320000 
--- FAIL: TestInt (0.00s)
panic: AmountDecimal与Decimal不相等
```



```go
var (
	intMin     int = math.MinInt
	intMax     int = intMin
	intNub     int     = 10000
	intDecimal int     = 20
)

panic: decimal division by 0 // Decimal返回error,直接panic
```

