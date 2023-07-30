[toc]

# 0. 前言

在阅读 `Kubernetes: kubectl` 源码时看到有关访问者设计模式的运用。访问者模式是行为型设计模式的一种，本篇文章将对访问者模式做一个介绍。

# 1. 访问者模式

## 1.1 示例
首先，给出一个比较粗糙的示例。  

当男人成功时，显示`我有一个好老婆`；当女人成功时，显示`我有一个有爱的丈夫`；
当男人开心时，显示`我有一个玩具`；当女人开心时，显示`我有一个有爱的丈夫`；
当男人伤心时，显示`我丢了玩具`；当女人伤心时，显示`我丢了有爱的丈夫`；

基于上述描述，实现示例代码：
```
type man struct {
	action string
}

type woman struct {
	action string
}

func (m *man) status() {
	if m.action == "happy" {
		fmt.Println("I have a toy")
	}

	if m.action == "sad" {
		fmt.Println("I lost my toy")
	}

	if m.action == "success" {
		fmt.Println("I have a great wife")
	}
}

func (w *woman) status() {
	if w.action == "happy" {
		fmt.Println("I have a lovely husband")
	}

	if w.action == "sad" {
		fmt.Println("I lost my lovely husband")
	}

	if w.action == "success" {
		fmt.Println("I have a lovely husband")
	}
}

func main() {
	m := man{
		action: "sad",
	}

	m.status()
}
```

示例代码实现了想要的功能。我们看示例代码的实现，男人和女人类都具有 `action` 且行为不一样。

那么，如果增加 `action`，比如当男人恋爱时，显示`我终于脱单了`；当女人恋爱时，显示`终于有人爱我了`；就需要更新男人和女人类的 `status` 方法，不符合开闭原则。

## 1.2 解耦对象和行为

既然不符合开闭原则，怎么样才能解耦对象和行为呢？使得行为的变化不会影响到对象实现。

还是进一步看上述实现，男人和女人类是行为的主体，行为有成功时的行为，开心时的行为，伤心时的行为...

既然是解耦，那先把行为拿出来作为接口，在以具体的行为类实现接口。看看怎么把行为类和对象类分开。

实现如下：
```
type status interface {
	manStatus()
	womanStatus()
}

type happy struct{}
type sad struct{}
type success struct{}

func (h *happy) manStatus() {
	fmt.Println("I have a toy")
}

func (s *sad) manStatus() {
	fmt.Println("I lost my toy")
}

func (s *success) manStatus() {
	fmt.Println("I have a great wife")
}

func (h *happy) womanStatus() {
	fmt.Println("I have a lovely husband")
}

func (s *sad) womanStatus() {
	fmt.Println("I lost my lovely husband")
}

func (s *success) womanStatus() {
	fmt.Println("I have a lovely husband")
}

type man struct {
}

type woman struct {
}
```

这里把对象类的 `action` 属性拿掉，相应的把 `action` 转换为行为类，行为类实现了接口其中定义了 `manStatus()` 和 `womanStatus()` 方法。

这样我们拆成了两个类，对象类和行为类。怎么把这两个类联系起来？
看代码，男人类和男人类的行为，女人类和女人类的行为是有联系的，初始化对象类和行为类：
```
m := man{}
s := sad{}

// m.demo(s)
```

应该通过 `demo` 将对象和行为类联系起来，然后调用 `s` 的 `manStatus()`方法完成关联：
```
func (m *man) accept(s status) {
	s.manStatus()
}

func (w *woman) accept(s status) {
	s.womanStatus()
}

func main() {
	m := &man{}
	s := &sad{}
	m.accept(s)
}
```

将 `demo` 重命名为 `accept` 作为对象类的方法，在其中调用行为类的方法。

## 1.3 访问者模式

至此，我们已经实现了一种新的设计模式访问者模式。有人可能会问了，我怎么还是没看出来。我们在分析上述代码。

有两个类，对象类和行为类。  
对象类实现 `accept` 方法，其是动作的主体，将对象类和行为类关联起来。  
行为类根据不同对象实现行为方法，其是行为的主体，行为是建立在对象之上的，也就是对于行为类来说，行为类是主。

基于上述分析，我们改造下代码如下：
```
type visitor interface {
	manStatus()
	womanStatus()
}

type happyVisitor struct{}
type sadVisitor struct{}
type successVisitor struct{}

func (h *happyVisitor) manStatus() {
	fmt.Println("I have a toy")
}

func (s *sadVisitor) manStatus() {
	fmt.Println("I lost my toy")
}

func (s *successVisitor) manStatus() {
	fmt.Println("I have a great wife")
}

func (h *happyVisitor) womanStatus() {
	fmt.Println("I have a lovely husband")
}

func (s *sadVisitor) womanStatus() {
	fmt.Println("I lost my lovely husband")
}

func (s *successVisitor) womanStatus() {
	fmt.Println("I have a lovely husband")
}

type man struct {
}

type woman struct {
}

func (m *man) accept(s visitor) {
	s.manStatus()
}

func (w *woman) accept(s visitor) {
	s.womanStatus()
}
```

结构基本没变，只是重命名了一下，更容易分清主体，站在不同类上，谁是主。

上述示例对于行为类相比于对象类是主的体验还不明显，重新改写代码：
```
type visitor interface {
	manStatus(*man)
	womanStatus(*woman)
}

type happyVisitor struct{}
type sadVisitor struct{}
type successVisitor struct{}

func (h *happyVisitor) manStatus(m *man) {
	fmt.Println(m.name, "I have a toy")
}

func (s *sadVisitor) manStatus(m *man) {
	fmt.Println(m.name, "I lost my toy")
}

func (s *successVisitor) manStatus(m *man) {
	fmt.Println(m.name, "I have a great wife")
}

type man struct {
	name string
}

func (m *man) accept(s visitor) {
	s.manStatus(m)
}

func main() {
	m := &man{"hxia"}
	s := &sadVisitor{}
	m.accept(s)
}
```

改写后的代码，行为类会访问对象类的 `name` 属性，对于行为类来说，对象类就是数据，是属性。

画出访问者设计模式的 UML 图：  
![访问者模式](image/访问者模式.png "访问者模式")

访问者模式在 GoF 合著的《设计模式：可复用面向对象软件的基础》中的定义是：
```
允许一个或多个操作应用到一组对象上，解耦操作和对象本身
Allows for one or more operation to be applied to a set of objects at runtime, decoupling the operations from the object structure
```

# 2. VisitorFunc 和访问者模式

前面介绍了访问者模式，从定义看出通过将一个或多个操作应用到一组对象上，以实现对象和操作的解耦。这里需要重点关注的点是一组对象。

一组意味着对象具有相似性，且结构是稳定的。试想如果男人和女人类中，女人没有伤心时的行为，那就没办法将其归为一组对象，或者需要实现 fake 伤心以保持对象的相似性。

既然定义的是一组，当然对象也可以是一个。假定对象是一个，使用函数改写上述访问者模式代码：
```
type VisitorFunc func(man)

func happyVisitor(m man) {
	fmt.Println(m.name, "I have a good thing")
}

func sadVisitor(m man) {
	fmt.Println(m.name, "I have a bad thing")
}

func successVisitor(m man) {
	fmt.Println(m.name, "I finish a thing")
}

type man struct {
	name string
}

func (m man) accept(v VisitorFunc) {
	v(m)
}

func main() {
	m := man{"hxia"}
	m.accept(sadVisitor)
}
```

这么改写在于简单，去掉类取而代之的是函数，用函数实现了具体的 `Visitor`。当然，这样的 `Visitor` 只能处理一种操作。

类似地，如果一组对象的行为是一样的，那也可以用函数 `Visitor` 来实现：
```
type VisitorFunc func(person)

func happyVisitor(p person) {
	fmt.Println(p.getName(), "I have a good thing")
}

func sadVisitor(p person) {
	fmt.Println(p.getName(), "I have a bad thing")
}

func successVisitor(p person) {
	fmt.Println(p.getName(), "I finish a thing")
}

type person interface {
	getName() string
}

type man struct {
	name string
}

type woman struct {
	name string
}

func (m man) getName() string {
	return m.name
}

func (w woman) getName() string {
	return w.name
}

func (m *man) accept(v VisitorFunc) {
	v(m)
}

func (w *woman) accept(v VisitorFunc) {
	v(w)
}

func main() {
	m := &man{"hxia"}
	m.accept(sadVisitor)
}
```

## 2.1 嵌套 Visitor

如果操作作用于一个对象，可以用函数 `Visitor` 来简化实现。如果多个操作嵌套的作用于对象上，那么可以使用嵌套 `Visitor` 实现，其效果类似于多个小应用访问数据库以实现某个功能。

举例如下：
```
type VisitorFunc func(*man) error

func happyVisitor(m *man) error {
	fmt.Println(m.name)
	return nil
}

func sadVisitor(m *man) error {
	fmt.Println(m.age)
	return nil
}

func successVisitor(m *man) error {
	fmt.Println(m.sex)
	return nil
}

func validationFunc(m *man) error {
	if m.name == "" {
		return errors.New("empty name")
	}

	return nil
}

type man struct {
	name string
	age  int
	sex  string
}

type Visitor interface {
	Visit(VisitorFunc) error
}

func (m *man) Visit(fn VisitorFunc) error {
	fmt.Println("in man")

	if err := fn(m); err != nil {
		return err
	}

	fmt.Println("out man")
	return nil
}

type validationVisitor struct {
	visitor Visitor
}

func (v validationVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in validation")

		if m.name == "" {
			return errors.New("empty name")
		}

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println("out validation")
		return nil
	})
}

type errorVisitor struct {
	visitor Visitor
}

func (v errorVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in error")

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println("out error")
		return nil
	})
}

type ageVisitor struct {
	visitor Visitor
}

func (v ageVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in age")

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println(m.name, m.age)

		fmt.Println("out age")
		return nil
	})
}

type VisitorList []Visitor

func (l VisitorList) Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(fn); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var visitor Visitor

	m1 := &man{name: "hxia", age: 18}
	m2 := &man{name: "huyun", age: 29}
	m3 := &man{name: "troy", age: 25}

	visitors := []Visitor{m1, m2, m3}

	visitor = VisitorList(visitors)
	visitor.Visit(happyVisitor)

	fmt.Println("================")

	visitor = validationVisitor{visitor: visitor}
	visitor.Visit(happyVisitor)

	fmt.Println("================")
	visitor = errorVisitor{visitor: visitor}
	visitor.Visit(happyVisitor)

	fmt.Println("================")

	visitor = ageVisitor{visitor: visitor}
	visitor.Visit(happyVisitor)
}
```

代码有点长，其基本是 `Kubernetes:kubectl` 访问者模式的主体，把它看懂了，再去看 `kubectl` 的访问者模式实现就不难了。

首先，对象实现了 `Visitor` (类似于上例的 accept)，接受函数 `Visitor`，函数 `Visitor` 访问对象操作：
```
func (m *man) Visit(fn VisitorFunc) error {
	fmt.Println("in man")

	if err := fn(m); err != nil {
		return err
	}

	fmt.Println("out man")
	return nil
}
```

接着，将多个对象装入 `VisitorList` 中，且该 VisitorList 也实现了 `Visitor` 接口方法。这么做是为了遍历访问每个对象：
```
func (l VisitorList) Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(fn); err != nil {
			return err
		}
	}
	return nil
}
```

然后，是在函数 `Visitor` 操作对象之后对对象做一些其它操作，这里定义了 `validationVisitor` 用来验证对象的名字是否为空：
```
func (v validationVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in validation")

		if m.name == "" {
			return errors.New("empty name")
		}

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println("out validation")
		return nil
	})
}

// main
visitor = validationVisitor{visitor: visitor}
visitor.Visit(happyVisitor)
```

通过层层嵌套 `Visitor` 实现了对对象的嵌套操作。

这些代码了解了，再去看 `Kubernetes:kubectl` 应该不难了，代码在[这里](https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/get/get.go)。
