# 说明

## 功能

将svn日志导出成excel

## 使用方法

- 修改配置文件`svnlog.yaml`为需要内容
	- 配置说明
	- `svnPath`,本地SVN目录(可以是svn的url),如:`E:\doc\Excel\China\0.8.0\Development\Excel`或`svn://dark.phoenix.wang/phx/Excel/China/0.8.0/Development/Excel`
	- `svnVersion`,需要导出的svn版本号,`from`与`to`的版本号先后顺序没有限制
	- `output`,输出文件的完整路径和文件名,因为是excel格式,扩展名需要为`xlsx`格式,如:E:\comment.xlsx
	- **yaml对空格有严格限制,注意不要随意删除空格**
- 运行`svnlog.exe`
- 生成的文件就是`output`指定的文件

## 例子

```yaml
# SVN日志导出工具
# 目录
svnPath: E:\doc\Excel\China\0.8.0\Development\Excel

# 版本号
svnVersion:
  from: 66096
  to: 66104

# 输出目录
output: E:\comment.xlsx
```
