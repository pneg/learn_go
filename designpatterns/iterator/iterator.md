## 概念
- 迭代器模式是一种行为设计模式，让你能在不暴露集合底层表现形式（列表、栈和树等）的情况下遍历集合中的所有元素
- 在迭代器的帮助下，客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素
- 这里需要注意的是有两个典型的迭代器接口，需要分清楚，一个是集合类实现的可以创建迭代器的工厂方法接口 一般命名为 Interable, 包含的方法类似 CreateIterator; 另一个是迭代器本身的接口，命名为 Iterator, 有Next 以及 hasMore 两个方法；
## 示例
- 一个班级类中包括一个老师和若干个学生，我们要对班级所有成员进行遍历，班级中老师存储在单独的结构字段中，学生存储再另外一个slice 字段中，通过迭代器，我们实现统一遍历处理。