学习笔记

业务服务分层处理

built-in error无法打印出堆栈信息，推荐使用pkg/errors

dao
捕获底层异常，进行wrap，向上传递

biz
没有异常，直接透传，有异常通过withMessage添加异常

service
统一处理异常：打日志，或返回错误

framework
通过Recovery兜底所有panic
