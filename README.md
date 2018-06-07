# pager
基于golang的一个分页类，当前仅设计了一种风格，其他的可以按照自己的需求添加。

# 样式如下:
  首页 上一页 1 2 3 4 5 6 7 下一页 尾页

# 使用方法
```
var allPage = 10    //总页数
var perNum = 7      //页码偏移量
var curPage = 2     //当前页码
var url = "/manage/msgList"
pager := NewPage(allPage,curPage,perNum,url)

pagelist := pager.OutPut()
```

把pagelist渲染到模版上即可.
