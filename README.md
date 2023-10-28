需求：
```
根据下文业务描述，使用Golang建模（无需完成 Repo层和 API 层，无需运行, 但是需要包含业务逻辑)。完成后提交代码到Github Public Repo.
回复邮件到 hr@pacport.com, 邮件名：应聘者姓名-Golang

TechSolutions is a national electronics retailer. They experienced rapid growth in the last year and have opened 50 new outlets. Each outlet sells electronic devices and related accessories, as well as store-specific gadgets. Outlets often have individual offers, but national marketing campaigns are often run, which influence the price of an item too.
TechSolutions recently launched a loyalty program called GadgetPoints, which allows customers to get 1 free accessory for every 10 they purchase. It doesn’t matter which outlet they purchase an item at or which they redeem it at.
TechSolutions has been thinking of launching an online store. They are also considering a monthly subscription that allows purchasers to get unlimited gadget servicing every month, as well as a discount on other devices. Now that we understand the business domain, we can start to explore how we can build systems to help TechSolutions achieve its goals!

TechSolutions是一家全国性的电子产品零售商。他们在去年经历了快速增长，新开了50家分店。每家商店都出售电子设备和相关配件，以及商店特有的小玩意。销售点通常有单独的优惠，但全国性的营销活动也会进行，这也会影响一件商品的价格。

TechSolutions最近推出了一项名为“GadgetPoints”的忠诚计划，该计划允许客户每购买10个配件即可获得1个免费配件。他们在哪个商店购买或在哪个商店兑换商品并不重要。

TechSolutions一直在考虑推出一家网上商店。他们还在考虑每月订阅，允许购买者每月获得无限的设备服务，以及其他设备的折扣。既然我们了解了业务领域，我们就可以开始探索如何构建系统来帮助TechSolutions实现其目标!
```

拆需求：
```
项目名：GadgetPoints
需求点：
1、全国性的电子产品零售商，新开了50家分店
2、出售电子设备、相关配件、商店特有的小玩意
3、销售点通常有单独的优惠，但全国性的营销活动也会进行
4、购买10个配件即可获得1个免费配件
5、设备的折扣

拆业务领域：
登录
注册
代理
产品分类
产品
活动
公告

接口(只提供了大概的)：
登录
1、登录：带上agentCode、username、password请求登录

注册
1、注册：带上agentCode、username、password等请求注册

代理
1、获取代理信息：带上agentCode、token请求获取代理信息

产品分类
1、获取产品分类：带上agentCode、token请求获取产品分类

产品
1、获取产品列表：带上agentCode、token、产品ID、请求获取产品

活动
1、获取活动列表：带上agentCode、请求活动列表

公告
1、获取公告列表：带上agentCode、请求公告列表

下单
1、下单：带上agentCode、token、产品ID、请求下单
2、获取订单列表：带上agentCode、token、获取订单列表
```