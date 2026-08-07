# 钱包商户对接文档

> 基础信息说明
---
1. 请求方式
请求方式：POST
数据格式：application/json

2. 签名须知
如何对请求参数进行签名，用户提交的参数除sign外，都要参与签名,
另外value为空的参数也不参与秘钥拼接,参数的key按照ASCII码值从小到大的排序
之后以key1=value1&key2=value2&...secret的方式拼接末尾拼接密钥
secret参数(只写值，不需要写变量名，不需要写'&')
最后，是利用32位md5算法，对最终待签名字符串进行签名运算，从而得到签名结果字符串(该字符串赋值于参数sign)。
样例

密钥 19aab5e6f1f4cb0ee3c2e5b4f86b780e2e95f2cf
签名字符串
amount=0.01&appKey=DmDpPP1B&currencyId=6&merchantOrderNo=W1767405902249328640&timestamp=171021737859719aab5e6f1f4cb0ee3c2e5b4f86b780e2e95f2cf
签名538dac4925529376b10d7e59204cfce4

3. 时间戳
时间戳都是13位毫秒级时间戳 

4. 金额
金额代表的是币的数量,小数位最低4位

5. 币种说明
- currencyId: 币种ID，默认为6（USDT-TRC20）
- itemId: 币种标识，USDT为10007
- itemName: 币种名称，使用小写，如：usdt
- chainTag: 链标识，使用小写，如：trc20、erc20等

6. 回调
回调, 第一次失败之后, 第二次为10秒之后, 第三次是 1 分钟之后

---

## 充值方式说明

商户对接支持两种充值方式：

### 1. 小程序充值
用户通过小程序页面完成充值，适合需要用户交互的场景
- 商户调用发起支付接口，获取支付链接
- 用户打开链接在小程序中完成支付
- 支付完成后系统回调商户

### 2. 地址充值
商户为用户创建专属充值地址，用户直接向地址转账即可
- 商户调用创建地址接口，为用户生成充值地址
- 用户向该地址转账
- 到账后系统回调商户
- 适合自动化充值、批量充值等场景

---

> Base domain: 

# 一、小程序充值

## 1.1 发起支付

> URI  /merchant/pay

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "merchantOrderNo": "w3493343423232",
  "attch": "",
  "remark": "string",
  "currencyId": 6,
  "amount": 0,
  "sign": "fe389787dd1f2ca5fa72407af919d489",
  "timestamp": 1710206777658
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||appKey|
|merchantOrderNo|string|true|none||商户订单号|
|attch|string|false|none||附加参数|
|remark|string|false|none||备注|
|currencyId|integer|true|none||币种id<br />币种id 6为usdt（TRC20链）|
|amount|number|true|none||金额,小数位4位(币的数量,指多少个)|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|

### 响应参数样例
```json
{
  "code": 0,
  "msg": "",
  "data": {
     "orderNo" : "W1770719928597348352",
     "currencyId" : 6,
     "userId" : 7,
     "tgUserId" : 5204224832,
     "amount" : "5",
     "orderStatus" : 0,
     "actualAmount" : "4.95",
     "merchantOrderNo" : "W1770719922763071488",
     "attch" : null,
     "remark" : null,
     "payTime" : null,
     "createTime" : 1711007504024,
     "timestamp" : 1711007504176,
     "url" : "https://t.me/find1WalletBot/mpay?startapp=merchantPayPage-W1770719928597348352"
  }
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer|false|none||id|
|orderNo|string|false|none||充值单号|
|currencyId|integer|false|none||币种id|
|amount|number|false|none||金额|
|orderStatus|integer|false|none||订单状态(0 待支付, 1 已支付, 2 已取消)|
|actualAmount|number|false|none||实际到账金额|
|merchantOrderNo|string|false|none||商户订单号|
|attch|string|false|none||附加参数|
|remark|string|false|none||备注|
|payTime| number |false|none||支付时间|
|createTime| number |false|none||下单时间|
|url|string|false|none||支付链接|




## 1.2 查询订单
> URI /merchant/queryOrder

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "orderNo": "www",
  "sign": "fe389787dd1f2ca5fa72407af919d489",
  "timestamp": 1710206777658
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||appId|
|orderNo|string|true|none||订单号|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|


### 响应参数样例

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    
    "orderNo": "W1770730458384957440", 
    "currencyId": 6, 
    "amount": "5", 
    "orderStatus": 1, 
    "actualAmount": "4.95", 
    "merchantOrderNo": "W1770730454895296512", 
    "attch": null, 
    "remark": null, 
    "payTime": 1711010029000, 
    "createTime": 1711010015000, 
    "timestamp": 1711010031127
  }
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer|false|none||id|
|orderNo|string|false|none||充值单号|
|currencyId|integer|false|none||币种id|
|amount|number|false|none||金额|
|orderStatus|integer|false|none||订单状态(0 待支付, 1 已支付, 2 已取消)|
|actualAmount|number|false|none||实际到账金额|
|merchantName|string|false|none||商户名称|
|merchantOrderNo|string|false|none||商户订单号|
|attch|string|false|none||附加参数|
|remark|string|false|none||备注|
|payTime| number |false|none||支付时间|
|createTime| number |false|none||下单时间|





## 1.3 回调通知

### 通知参数样例
```json
{
    "actualAmount": 4.95, 
    "amount": 5, 
    "createTime": 1711008030000, 
    "currencyId": 6, 
    "merchantOrderNo": "W1770722135782719488", 
    "orderNo": "W1770722136076320768", 
    "orderStatus": 1, 
    "payTime": 1711008147153, 
    "sign": "ca984c96933d320f6a74efaaf7ba4a67", 
    "tgUserId": 5204224832, 
    "attch": "", 
    "remark": "", 
    "timestamp": 1711008158182, 
    "userId": 7, 
    "appKey": "DmDpPP1B"
}
```
### 回调返回
```String
OK
```
### 参数说明

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appkey|string|false|none||商户key|
|sign|string|false|none||签名|
|timestamp|integer|false|none||时间戳<br />当前时间戳13位|
|orderNo|string|false|none||充值单号|
|currencyId|integer|false|none||币种id|
|userId|integer|false|none||用户id|
|tgUserId|integer|false|none||tg用户Id|
|amount|number|false|none||金额|
|orderStatus|integer|false|none||订单状态(0 待支付, 1 已支付, 2 已取消)|
|actualAmount|number|false|none||实际到账金额|
|merchantOrderNo|string|false|none||商户订单号|
|attch|string|false|none||附加参数|
|remark|string|false|none||备注|
|payTime|string|false|none||支付时间|
|createTime| number |false|none||下单时间|

---

# 二、地址充值

## 2.1 创建Mask地址

> URI  /merchant/createAddress

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "userId": 123456,
  "itemId": 10007,
  "itemName": "usdt",
  "chainTag": "trc20",
  "timestamp": 1710217378597,
  "sign": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||商户appKey|
|userId|integer|true|none||用户ID|
|itemId|integer|true|none||币种id<br />币种id 10007为usdt|
|itemName|string|false|none||币种名称|
|chainTag|string|true|none||链标识<br />如：trc20、erc20等（小写）|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|

### 响应参数样例
```json
{
  "code": 0,
  "msg": "success",
  "data": "TXxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|code|integer|false|none||响应码 0表示成功|
|msg|string|false|none||响应消息|
|data|string|false|none||创建的充值地址|

### 注意事项
1. **IP白名单**: 请求IP必须在商户配置的白名单中（如果启用IP检查）
2. **幂等性**: 相同的userId、itemId、chainTag会返回相同的地址
3. **地址复用**: 如果该用户已经创建过该币种和链的地址，会直接返回已有地址

---

## 2.2 查询充值订单

> URI  /merchant/queryRecharge

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "upOrderNo": "A1770719928597348352",
  "clientUserId": "123456",
  "itemId": 10007,
  "itemName": "usdt",
  "chainTag": "trc20",
  "timestamp": 1710217378597,
  "sign": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
}
```

### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||商户appKey|
|upOrderNo|string|true|none||充值订单号（回调中的upOrderNo）|
|clientUserId|string|true|none||商户用户ID|
|itemId|integer|false|none||币种id|
|itemName|string|false|none||币种名称|
|chainTag|string|false|none||链标识|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|

### 响应参数样例
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "amount": 100.0000,
    "itemId": 10007,
    "chainTag": "trc20",
    "txid": "0x1234567890abcdef...",
    "fromAddress": "TXxxxxxxxxxxxxxxxxxx",
    "toAddress": "TYxxxxxxxxxxxxxxxxxx",
    "upOrderNo": "A1770719928597348352",
    "createTime": 1710217378597
  }
}
```

### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|code|integer|false|none||响应码 0表示成功|
|msg|string|false|none||响应消息|
|data|object|false|none||订单信息|
|amount|number|false|none||充值金额|
|itemId|integer|false|none||币种ID|
|chainTag|string|false|none||链标识|
|txid|string|false|none||交易hash|
|fromAddress|string|false|none||来源地址|
|toAddress|string|false|none||目标地址（充值地址）|
|upOrderNo|string|false|none||充值订单号|
|createTime|integer|false|none||订单创建时间<br />13位时间戳|

---

## 2.3 充值回调通知

当用户向创建的地址充值成功后，系统会自动回调商户配置的充值回调地址

> 回调地址: 商户后台配置的 notifyUrl

### 通知参数样例
```json
{
  "appKey": "DmDpPP1B",
  "currencyId": 6,
  "timestamp": 1710217378597,
  "upOrderNo": "A1770719928597348352",
  "amount": 100.0000,
  "clientUserId": "123456",
  "itemId": 10007,
  "chainTag": "trc20",
  "sign": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "status": "success",
  "orderType": 1,
  "createTime": 1710217378597
}
```

### 回调返回
```String
OK
```

### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|false|none||商户appKey|
|currencyId|integer|false|none||币种id|
|timestamp|integer|false|none||时间戳<br />当前时间戳13位|
|upOrderNo|string|false|none||充值订单号|
|amount|number|false|none||充值金额|
|clientUserId|string|false|none||商户用户ID（创建地址时传入的userId）|
|itemId|integer|false|none||币种ID|
|chainTag|string|false|none||链标识|
|sign|string|false|none||签名|
|status|string|false|none||订单状态<br />success成功 faild失败|
|orderType|integer|false|none||订单类型<br />1充币 2提币|
|createTime|integer|false|none||订单创建时间<br />13位时间戳|

---

# 三、小程序提现（商户提现到用户）

## 3.1 发起提现

> URI  /merchant/sendMoney

商户将余额提现到指定TG用户，通过tgUserId指定收款用户，如果用户不存在会自动创建

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "orderNo": "W1770719922763071488",
  "tgUserId": 5204224832,
  "userTgName": "张三",
  "userTgUsername": "zhangsan",
  "amount": 5.0000,
  "currencyId": 6,
  "sign": "fe389787dd1f2ca5fa72407af919d489",
  "timestamp": 1710206777658
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||商户appKey|
|orderNo|string|true|none||商户订单号|
|tgUserId|integer|true|none||收款用户TG ID|
|userTgName|string|false|none||收款用户TG昵称|
|userTgUsername|string|false|none||收款用户TG用户名|
|amount|number|true|none||金额,小数位4位(币的数量,指多少个)|
|currencyId|integer|true|none||币种id<br />币种id 6为usdt|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|

### 响应参数样例
```json
{
  "code": 0,
  "msg": "",
  "data": {
    "orderNo": "W1770719928597348352",
    "currencyId": 6,
    "amount": 5.0000,
    "createTime": 1711007504024,
    "timestamp": 1711007504176
  }
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|orderNo|string|false|none||提现单号|
|currencyId|integer|false|none||币种id|
|amount|number|false|none||金额|
|createTime|number|false|none||提现时间|
|timestamp|number|false|none||时间戳|

### 注意事项
1. 该接口无回调通知，提现实时到账
2. 相同的appKey + orderNo具有幂等性，重复请求会返回已有订单信息
3. 如果tgUserId对应的用户不存在，系统会自动创建用户
4. 提现金额从商户余额中扣除，会收取手续费

---

# 四、地址提现（冷钱包提现）

## 4.1 发起提现

> URI  /merchant/goldWalletWithdraw

商户将余额提现到指定区块链地址，提现金额超过审核阈值时需要人工审核

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "orderNo": "W1770719922763071488",
  "address": "TXxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "amount": 100.0000,
  "currencyId": 6,
  "sign": "fe389787dd1f2ca5fa72407af919d489",
  "timestamp": 1710206777658
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||商户appKey|
|orderNo|string|true|none||商户订单号|
|address|string|true|none||提现地址（区块链地址）|
|amount|number|true|none||金额,小数位4位(币的数量,指多少个)|
|currencyId|integer|true|none||币种id<br />币种id 6为usdt|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|

### 响应参数样例
```json
{
  "code": 0,
  "msg": "",
  "data": {
    "orderNo": "L1770719928597348352",
    "merchantOrderNo": "W1770719922763071488",
    "currencyId": 6,
    "amount": 100.0000,
    "actualAmount": 99.0000,
    "fee": 1.0000,
    "createTime": 1711007504024,
    "timestamp": 1711007504176
  }
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|orderNo|string|false|none||提现单号|
|merchantOrderNo|string|false|none||商户订单号|
|currencyId|integer|false|none||币种id|
|amount|number|false|none||金额|
|actualAmount|number|false|none||实际到账金额|
|fee|number|false|none||手续费|
|createTime|number|false|none||提现时间|
|timestamp|number|false|none||时间戳|

### 注意事项
1. 相同的appKey + orderNo具有幂等性，重复请求会返回已有订单信息
2. 提现金额需大于等于币种最小提现金额
3. 提现金额超过商户审核阈值时，订单进入待审核状态，需人工审核通过后才会发起链上转账


## 4.2 查询提现订单

> URI  /merchant/queryWithdrawOrder

### 请求参数样例
```json
{
  "appKey": "DmDpPP1B",
  "orderNo": "L1770719928597348352",
  "sign": "fe389787dd1f2ca5fa72407af919d489",
  "timestamp": 1710206777658
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|true|none||商户appKey|
|orderNo|string|true|none||订单号<br />支持传入系统订单号或商户订单号，两者均可查询|
|sign|string|true|none||签名|
|timestamp|integer|true|none||时间戳<br />当前时间戳毫秒值13位|

### 响应参数样例
```json
{
  "code": 0,
  "msg": "",
  "data": {
    "orderNo": "L1770719928597348352",
    "currencyId": 6,
    "amount": 100.0000,
    "orderStatus": 4,
    "actualAmount": 99.0000,
    "merchantOrderNo": "W1770719922763071488",
    "remark": "",
    "createTime": 1711007504024,
    "timestamp": 1711007504176,
    "sign": "ca984c96933d320f6a74efaaf7ba4a67"
  }
}
```
### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|orderNo|string|false|none||提现单号|
|currencyId|integer|false|none||币种id|
|amount|number|false|none||金额|
|orderStatus|integer|false|none||订单状态<br />3 已拒绝, 4 提现成功, 5 提现失败, 其他状态为处理中|
|actualAmount|number|false|none||实际到账金额|
|merchantOrderNo|string|false|none||商户订单号|
|remark|string|false|none||备注|
|createTime|number|false|none||下单时间|
|timestamp|number|false|none||时间戳|
|sign|string|false|none||签名|


## 4.3 提现回调通知

提现完成后（成功或失败），系统会自动回调商户配置的提现回调地址

> 回调地址: 商户后台配置的 withdrawNotifyUrl

### 通知参数样例
```json
{
  "appKey": "DmDpPP1B",
  "currencyId": 6,
  "timestamp": 1710217378597,
  "upOrderNo": "L1770719928597348352",
  "merchantOrderNo": "W1770719922763071488",
  "amount": 100.0000,
  "actualAmount": 99.0000,
  "fee": 1.0000,
  "itemId": 10007,
  "chainTag": "trc20",
  "sign": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
  "status": "success",
  "orderType": 2,
  "createTime": 1710217378597
}
```

### 回调返回
```String
OK
```

### 参数说明
|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|appKey|string|false|none||商户appKey|
|currencyId|integer|false|none||币种id|
|timestamp|integer|false|none||时间戳<br />当前时间戳13位|
|upOrderNo|string|false|none||提现订单号|
|merchantOrderNo|string|false|none||商户订单号|
|amount|number|false|none||提现金额|
|actualAmount|number|false|none||实际到账金额|
|fee|number|false|none||手续费|
|itemId|integer|false|none||币种ID|
|chainTag|string|false|none||链标识|
|sign|string|false|none||签名|
|status|string|false|none||订单状态<br />success成功 faild失败|
|orderType|integer|false|none||订单类型<br />1充币 2提币|
|createTime|integer|false|none||订单创建时间<br />13位时间戳|
