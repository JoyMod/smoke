package smoke

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 构造一个MAP来接收封装的内容
type H map[string]interface{}

// Context 构造
type Context struct {
	//封装请求和响应
	Writer http.ResponseWriter
	Req    *http.Request

	//封装配置
	Path   string
	Method string
	Params map[string]string

	//封装响应码
	StatusCode int

	//中间件
	handlers []HandlerFunc
	index    int

	engine *Engine
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// newContext 初始化
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// PostForm 表单抓取的方法
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query String 抓取方法
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 网页状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置方法
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 字符串配置
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 配置
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 配置
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 配置
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}

func (c *Context) Fail(s int, p string) {
	s = c.StatusCode
	p = ""
}
