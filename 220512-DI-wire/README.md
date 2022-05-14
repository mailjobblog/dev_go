# Go依赖注入工具wire

## 介绍

Go 语言常用的依赖注入工具有 google/wire、uber-go/dig、facebookgo/inject，我们以 wire 和 dig 作比较，对比一下两个工具的区别。

- dig 通过反射识别依赖关系，wire 是编译前计算依赖关系，Wire 作为代码生成器运行，这意味着注入器无需调用运行时库即可工作。
- dig 只能在代码运行时，才能知道哪个依赖不对，比如构造函数返回类型的是结构体指针，但是其他依赖的是interface，这样的错误只能在运行时发现，而wire可以在编译的时候就发现。
- 由于采用了依赖注入，所以在代码调试时可以注入一些mock 服务或者函数，wire在mock上支持更友好些，dig的话可以通过build tag 来使用mock。 个人比较推荐使用wire，可以在编译时就发现问题，避免了 多次的build和尝试后才解决编译问题

## 官方文档

- wire github：https://github.com/google/wire

## Wire详解

本文基于 `wire v0.5.0` 编写，关于测试代码可在 [DI-wire](https://github.com/mailjobblog/dev_go/tree/master/220512-DI-wire) 下载。

### 搭建测试服务

在测试 wire 过程中，可能需要用到一些 DB 服务，请根据 example 的需求，安装需要的服务。

```bash
docker run -itd -p 3310:3306 -e MYSQL_ROOT_PASSWORD=root --name wire-mysql mysql:5.7
```



## 参考资料

**wire**

- https://segmentfault.com/a/1190000039185137
- **https://juejin.cn/post/6844903901469097998**
- **https://www.modb.pro/db/231728**
- https://darjun.github.io/2020/03/02/godailylib/wire
- https://www.cnblogs.com/Me1onRind/p/13624487.html


**wire工程化**

- https://github.com/go-kratos/beer-shop