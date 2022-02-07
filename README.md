# mylorca
golang 基于 lorca gui架构



### 创建UI

```go
view.New(800, 600)

view.UI.Wait()
```

### 创建控制器

```
view.New(800, 600)
ctl.NewIndexCtl()
view.UI.Wait()
```

### 控制器定制

```
//继承Ctl结构体（自动绑定ui事件）
type IndexCtl struct {
	Ctl
	input *element.Input
	sp1   *element.Span
}
//初始化并与页面ui元素通过id进行绑定
func NewIndexCtl() *IndexCtl {
	c := &IndexCtl{}
	c.input = &element.Input{Id: "input1"}
	c.sp1 = &element.Span{Id: "sp1"}
	c.bind(c)
	return c
}
//定义方法绑定（方法名与前端调用事件名相同）
func (c *IndexCtl) Hello() {
	c.input.SetValue(c.input.GetValue() + "s")
	c.sp1.SetValue(c.input.GetValue())
}
```

### 组件定义

```
//创建组件结构体
type Input struct {
	Id    string
	value string
}
//定义组件与js关联支持方法
func (m *Input) SetValue(value string) {
	js := fmt.Sprintf(setValueStr, m.Id, value)
	view.UI.Eval(js)
}
//定义组件与js关联支持方法
func (m *Input) GetValue() string {
	js := fmt.Sprintf(getValueStr, m.Id)
	value := view.UI.Eval(js)
	m.value = value.String()
	return m.value
}

```

### html+bootstrup定义页面（仅使用bootstrap 样式文件）

```
<!doctype html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>主页</title>
    <link rel="stylesheet" href="./stc/bootstrap.min.css"></link>
    
  </head>
  <body>
    <span id="sp1">2222222</span>
    <div class="form-group">
      <label for="exampleInputEmail1">Email address</label>
      <input type="email" class="form-control" id="input1" placeholder="Email">
    </div>
    <button type="submit" onclick="Hello()"  class="btn btn-default">Submit</button>
  </body>
</html>
```

