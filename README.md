# goframework-resty

| 名称                     |方法| 描述          |
|------------------------| --- |-------------|
| 获取 `*resty.Client`     | GetRestyClient |             |
| 获取 `*goresty.Feign`    | GetFeignClient | 不存在时会按需创建默认实例 |
| 通过 `*resty.Client`添加实例 | AddFeignInstance              ||
| 通过 `*http.Client`添加实例  | AddFeignInstanceWithClient    ||
| 通过 `net.Addr`添加实例      | AddFeignInstanceWithLocalAddr | 使用本地地址绑定 |
| 移除实例                   | RemoveFeignInstance     ||
| 实例是否存在                 | HasFeignInstance        ||
| 配置回调                   | Options                 | 配置参数以回调方式配置 |
| 错误回调                   | OnError                 ||
| 发送拦截器                  | OnBeforeRequest         ||
| 响应拦截器                  | OnAfterResponse         ||
| 请求对象                   | Request                 | 获取 *resty.Request |
| 请求对象（带 Context）       | RequestWithContext      | 便于注入上下文       |

- 方法列表

```go
func GetRestyClient(name string) *resty.Client
func GetFeignClient(name string) *goresty.Feign
// AddFeignInstance 添加feign
func AddFeignInstance(name string, r *resty.Client) error
// AddFeignInstanceWithClient 添加feign
func AddFeignInstanceWithClient(name string, hc *http.Client) error
// AddFeignInstanceWithLocalAddr 添加feign
func AddFeignInstanceWithLocalAddr(name string, localAddr net.Addr) error
// RemoveFeignInstance 移除feign
func RemoveFeignInstance(name string)
// HasFeignInstance feign句柄是否存在
func HasFeignInstance(name string) bool
func Options(name string, f func(*resty.Client))
func OnError(name string, h resty.ErrorHook)
func OnBeforeRequest(name string, m resty.RequestMiddleware)
func OnAfterResponse(name string, m resty.ResponseMiddleware)
func Request(name string) (*resty.Request, error)
func RequestWithContext(name string, ctx context.Context) (*resty.Request, error)
```

## 使用示例

```go
// 初始化
// 若未预先添加，第一次 GetFeignClient/GetRestyClient 会自动创建默认实例
_ = AddFeignInstance("default", resty.New())
Options("default", func(c *resty.Client) {
    c.SetBaseURL("https://example.com")
})

// 普通请求
req, _ := Request("default")
resp, err := req.SetResult(&out).Get("/api/ping")

// 带 Context
ctx := context.WithValue(context.Background(), "k", "v")
req2, _ := RequestWithContext("default", ctx)
resp2, err := req2.SetResult(&out).Get("/api/ping")
_ = resp2
_ = err
```
