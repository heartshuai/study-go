package coding

/**
为什么要代码规范
	1.代码规范并不是强制的 但是不同的语言 一些细微的规范还是要遵守的
	2.代码规范主要是为方便团队内部行程一个统一的代码风格，提高代码的可读性，统一性
1.代码规范
	1.命名规范
		包名
			尽量和目录保持一致
			尽量采取有意义的包名尽量的简短
			包名不要和标准库名冲突
			包名全部小写
		文件名
			user_name 如果有多个单词可以采用蛇形命名法
		变量名
			1.蛇形命名法 python php
			2.驼峰命名法 java c go
				userName
				un string
				有一些专有名词 全部采用大写或者小写 APIVersion
				bool类型 Has Is can allow 开头
		结构体命名
			驼峰  UserName
		接口命名
			基本上和结构体差不多
			接口以er结尾
			type Writer interface{}
			type IRead	interface{}
		常量命名
			 全大写 如果有多个单词 使用蛇形命名法  APP_VERSION
2.注释规范
	1.go提供两种注释
 		1.//单行注释
		2.适合大段的注释
			变量注释
			包注释
			函数注释
			接口注释
			代码逻辑注释


3.import规范

4.错误处理
*/
