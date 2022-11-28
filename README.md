# 错误码定义

|  错误码  | 错误信息            |
|:-----:|:----------------|
|   0   | 成功              |
| 10000 | A类错误: 用户类错误     |
| 20000 | B类错误: 参数错误      |
| 30000 | C类错误: 第三方服务调用错误 |
| 31000 | C类错误: 数据库错误     |
| 32000 | C类错误: redis错误   |
| 40000 | D类错误: 服务内部错误    |
| 99999 | 未明确的错误          |

## 错误码使用示例
- 统一使用错误码 错误信息两个字段
- 处理函数handler返回的error是标准库error，所以需要使用errs模块生成error，否则返回未知错误码
- demo如下：
```golang
import "gitlab.ctyuncdn.cn/wangzhk/errs"

func Init() {
	errs.Init(errs.EsportsSrv)
}

func InternalCall() *errs.Error {
	// .....................
	return errs.ErrInfo[errs.AUserErr] // 推荐
}


func (s *GreeterServerImpl) SayHello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloReply) (err error) {
    //...........
	err := InternalCall(xxxx)
	
    return errs.I18nTrans(ctx, err)
}
```
