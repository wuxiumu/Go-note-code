# HTTP请求中的referrer和Referrer-Policy
本文将介绍一个涉及安全和隐私的http请求头中的字段—referrer，以及如何通过Referrer Policy去修改referrer的值或者是显示与否。

## 什么是referrer
当一个用户点击当前页面中的一个链接，然后跳转到目标页面时，目标页面会收到一个信息，即用户是从哪个源链接跳转过来的。如下图所示：

![](/img/16aac3dfed455faa.png)

也就是说，当你发起一个http请求，请求头中的referrer字段就说明了你是从哪个页面发起该请求的。

# 使用场景
有时候我们需要控制这个referrer字段的值，即是否让其显示在请求头中，或者是否显示完整路径等。尤其是在以下两个使用场景：

## 隐私
在社交网站的个人中心页面，也许会存在一些外链，这时候社交网站肯定不希望用户在点击这些链接跳转到其他第三方网站时会将自己个人中心的URL信息显示在referrer字段中传过去，尤其是个人中心页面的URL往往会带着用户数据和一些敏感信息。这时候可以选择不显示来源页面URL信息或者只显示一个网站根地址hostname。

## 安全
有些使用了https的网站，可能在URL中使用一个参数（sid）来作为用户身份凭证，而又需要引入其他https网站的资源，这种情况，网站肯定不希望泄露用户的身份凭证信息。当https网站需要引入不安全的http网站的资源或者有链接要跳转到http网站时，这时候将https源网站的URL信息传过去也是不太安全的。

当然还有其他情况下需要referrer的值，比如最近公司所做的项目中，有一个请求由于请求头过大导致响应是400，我们的Referrer Policy是默认的情况，显示的referrer是完整的URL信息，该URL带了很多敏感数据比如加密后的token，sessionID等，长度特别长，请求头中的cookie和请求的URL也带着很大块的信息，最终我们决定让referrer只携带网站根地址的信息而不是其完整路径，由此减小了header的大小。

## Referrer-Policy
Referrer-Policy的作用就是为了控制请求头中referrer的内容，目前是一个候选标准，不过已经有部分浏览器支持该标准。
目前Referrer-Policy只包含以下几种值：
```
enum ReferrerPolicy {
    "",
    "no-referrer",
    "no-referrer-when-downgrade",
    "same-origin",
    "origin",
    "strict-origin",
    "origin-when-cross-origin",
    "strict-origin-when-cross-origin",
    "unsafe-url"
};
```

## 空字符串
若设为空串则默认按照浏览器的机制设置referrer的内容，默认情况下是和no-referrer-when-downgrade设置得一样。

## no-referrer
不显示referrer的任何信息在请求头中。

## no-referrer-when-downgrade
这是默认值。当从https网站跳转到http网站或者请求其资源时（安全降级HTTPS→HTTP），不显示referrer的信息，其他情况（安全同级HTTPS→HTTPS，或者HTTP→HTTP）则在referrer中显示完整的源网站的URL信息。

## same-origin
表示浏览器只会显示referrer信息给同源网站，并且是完整的URL信息。所谓同源网站，是协议、域名、端口都相同的网站。

## origin
表示浏览器在referrer字段中只显示源网站的源地址（即协议、域名、端口），而不包括完整的路径。

## strict-origin
该策略更为安全些，和origin策略相似，只是不允许referrer信息显示在从https网站到http网站的请求中（安全降级）。

## origin-when-cross-origin
当发请求给同源网站时，浏览器会在referrer中显示完整的URL信息，发个非同源网站时，则只显示源地址（协议、域名、端口）

## strict-origin-when-cross-origin
和origin-when-cross-origin相似，只是不允许referrer信息显示在从https网站到http网站的请求中（安全降级）。

## unsaft-url
浏览器总是会将完整的URL信息显示在referrer字段中，无论请求发给任何网站。

# Referrer-Policy更改方法
可以有以下5种方法：
## 1. 通过Referrer-Policy HTTP header设置：
```
Referrer-Policy: origin
```

## 2. 通过<meta>元素改变Referrer Policy，直接修改名为referrer的内容
```
<meta name="referrer" content="origin">
```

 
## 3. 给
```
<a>, <area>, <img>, <iframe>, 
```
或者
```
<link>
```
元素设置referrerpolicy属性
 
```
<a href="http://example.com" referrerpolicy="origin">
```

## 4. 如需设置不显示referrer信息时，也可以给 
```
<a>, <area>, <link>
```
元素设置rel的链接关系。

```
<a href="http://example.com" rel="noreferrer">
```

# 总结
使用何种Referrer Policy取决于网站的需求，但是一般来说，unsafe-url是不太建议用的，同样，如果是只想显示网站的根地址，那么建议用strict-origin和strict-origin-when-cross-origin。如果URL中没有什么敏感信息，那就默认使用no-referrer-when-downgrade。



参考：https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
