# pingyin

内存中的汉字和拼音对照

使用 PH, HP 两个变量
PH : 拼音 -> 汉字
HP : 汉字 -> 拼音

```go
func Test_1(t *testing.T) {
	assert.Equal(t, []string{"wo"}, pingyin.HP["我"])
	assert.Equal(t, []string{"ai"}, pingyin.HP["爱"])
	assert.Equal(t, []string{"ni"}, pingyin.HP["你"])
}

func Test_2(t *testing.T) {
	assert.Equal(t, []string{"hang", "xing"}, pingyin.HP["行"])
}

func Test_3(t *testing.T) {
	t.Logf("%v\n", pingyin.PH["ya"])
}
```