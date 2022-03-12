#  策略模式

定义一个接口统一所有相似算法，定义新的结构体具体实现这一接口，增加Context结构维护对Strategy具体对象的引用，与工厂模式不同的是，策略模式使用于多个相似算法之间的转化，而工厂模式解决的是多个对象实例化问题。同时策略也符合开闭原则。