# 超级管理员登入/登出 功能模块 #
 
---

相关数据表模型结构 SysUser

定义文件: migrate/check.go



**数据库表结构**

| 字段 | 定义 | 描述 |
| ---- | ---- | ----| 
|id|mint(8)	unsigned auto_increment primary|自增ID|
|uname|varchar(30) not null unique|管理员登录名|
|upass|varchar(32) not null|管理员HASH后密码(util.SysPass后)|
|Ustat|int unsigned not null,default:1,|用户状态|
|CreatedAt| | GORM自带|
|UpdatedAt| | GORM自带|
|DeletedAt|index | GORM自带|

---

功能需求部分，请实现后并测试无误后，用实际简单文档插入到这里，将需求文档逐渐压到最后，完成全部测试后，删除需求文档。

**功能需求列表：**

1. 超级后台登入(区别于用户登入)
2. 超级后台登出
3. 超级用户修改本人密码

**详情：**

- 用户登入

		URI: /superadm/login
		METHOD: POST
		Format: JSON
			username: (string)用户名明文
			password: (string)前端md5后的用户输入密码，过滤首尾空格

- 返回

		Format: JSON
		
		登录成功:
				result: (int)1
		登录失败:
				result: (int)0
				reason: (string)失败详情
 
		
- 服务器端状态

		Session: 
				生成并存入以下信息	: ID,Uname

- 前端操作

		登录成功:
				Redirect: 跳转至超级用户后台首页静态页面
		登录失败:
				Alert提示后清空用户密码框，让用户重新填写


- 后端逻辑提示

		1. 判断登录应按用户名取回记录，采用util.SysPass方式HASH后进行比对
		2. 所有前后端传输密码均应在前端md5一次
