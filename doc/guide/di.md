# 依赖注入

依赖注入是保持软件`低耦合、易维护`的重要设计准则之一。 此准则被广泛应用在各种开发平台之中，有很多与之相关的优秀工具。 其中最著名的当属`Spring`，`Spring IOC`作为框架的核心功能对`Spring`
的发展到今天统治地位起了决定性作用。 事实上，软件开发`S.O.L.I.D` 原则 中的`D`， 就专门指代这个话题

## 盘古的依赖注入

实际上Golang里面的依赖注入还有争议（有一部分Golang开发者认为Golang不需要依赖注入），在这里我引入在网上看到的一个经典回答来回复这个问题
> 像Google、Uber以及Facebook这种大公司都在使用依赖注入，你凭什么不使用

事实上，依赖注入的最大好处是
> 在系统变得复杂的时候，依然保持系统简单

## 核心概念

也许有的童鞋对依赖注入已经很熟悉（尤其是从Java语言转Golang语言的童鞋），但是从我的观测来看，大部分Gopher对依赖注入完全没有概念，所以花一点时间来解释依赖注入的核心概念，有助于更了解`盘古`应用程序框架

### 构造函数

Golang里面没有`构造函数`这个说法，但是并不妨碍我们写`构造函数`，简单的说，`构造函数`就是返回某种类型的简单方法

<<< @/../example/guide/constructor.go

需要注意的是

- 构造函数并不要求返回指针类型
- 构造函数可以是公开的，也可以私有的（写成私有的能最大限度的隐藏实现，当然这需要在代码编写过程中去平衡）

### 添加依赖

可以添加任何功能实现方到系统内部，然后调用方就能在任何地方使用这些功能了

<<< @/../example/guide/pangu.go

建议

- 最好是用`Struct`封装功能
- 尽量私有优先，隐藏内部实现
- 并不要求返回指针
- 使用`init`和`Musts`在出现依赖关系出错时会导致应用启动时`panic`，可以及时发现出错的依赖
- `pangu.New()`是单例且线程安全的，可以在任何地方放心使用

### 使用依赖

在`盘古`里面使用依赖非常简单，声明需要使用的依赖即可

<<< @/../example/guide/server.go

- 需要什么依赖，参加到方法参数里面即可
- `盘古`会在整个系统中找到依赖项，如果没有找到，会抛出`error`，应用程序可以自行处理错误

### 组合依赖

组合依赖的原因是因为有可能依赖项太多，都写到方法参数里面显得太杂乱，要组合依赖关系特别的容易，只需要继承`pangu.In`就可以了

<<< @/../example/guide/in.go

注意

- 可以组合的依赖无限制
- 组合的依赖必须是可导出的（公开）
- 组合依赖包装结构体不能是指针类型
