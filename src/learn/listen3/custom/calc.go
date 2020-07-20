package custom

var (
	Sum int
)

/**
自定义包  执行go install 后在在pkg中生成.a文件
*/
func TowSum(a, b int) int {
	return a + b
}
