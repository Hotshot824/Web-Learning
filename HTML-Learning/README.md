# HTML

**學習關於HTML/CSS/JavaScript筆記**

**參考清大Web-Programming開放式課程**

* How does internet Work?
* Web Development

<br>

# How does internet Work?

**What Happened?**

當在URL按下enter，將發生什麼？

https://en.wikipedia.org/wiki/Main_Page

![](https://i.imgur.com/OjHk8dz.png)

* `https://` : **Protocol** 使用何種通訊協定

* `en.wikipedia.org` : **Domain name** 連線至哪一個網域

* `/wiki/Main_Page` : **Path** 網域所能提供的 Resources

<br>

首先將Domain Name 拿出詢問 `DNS` 這個Domain Name所代表的 `IP(Internet Protocol)`

`en.wikipedia.org` 是給 User 使用記憶，如`www.wikipedia`代表不同語言的服務。

而 Machine 所使用的則是 IP 之後 Browser 才可以連到對應的 Server。

這時便可以與 Server 建立 Connection 可能是 `UDP` 或 `TCP`。  

<br>
 
![](https://i.imgur.com/2ZXNkNT.png)

當與 Server 連線後便可以使用 HTTP `request` 如 GET/documentart.mpe

Server `Response` 200 OK 與 Page，這樣就可以在 Browser 上呈現 Page 內容。

這就是按下 Enter 之後到 Browser 上呈現背後所經過的過程。

<br>

## HTTP Protocol

* A `protocol` is language spoken by machines
  * Defines structure of messages to be exchanged
* HyperText Transfer Protocol(HTTP) defines:
  * Messages : HTTP `request` and HTPP `response`
  * Requests : accessing `resources` (web pages) via `GET POST PUT DELETE` methods
  * Responses : **2**00 OK, **3**01 Moved, **4**04 Not Found, **5**00 Server Error,etc.

關於 [Requests](https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Methods) 與 [Responses](https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status) 詳細的說明

<br>

## HTTP Messages

![](https://i.imgur.com/hyVrpll.png)

* Each HTTP message have
  * `start line`, `header line`, and optionally `body`
* A resource can have different `presentations`
  * HTML, XML, JSON, etc.

從上面的 Example 可以看到 `start` 中 Requests GET 與 Protocol 的版本名稱，Server Respones 200 OK，`Body` 則是交換的 Resources, 這些交換的 Resources 並不限定於網頁。

Requests `Header`，包含自己的 `Host`，`Accept`，`Accept-Encoding` 等代表著接受的端口、檔案格式等資訊

Response `Header`，則有 `Content-Type`，`Content-Length` 代表 Body 的格式，長度等資訊

<br>

## Chrome inspect Example

![](https://i.imgur.com/E19kX8K.jpg) 

從 `Chrome inspect` 上我們就可以看到 HTTP 交換的 Header，在這裡從 www.wikipedia.org 送出之後 Response 301 因此 Browser 就轉向尋找 Location，然後就轉向 en.wikipedia.org 

第二次的 Requests 得到 200 的 Response，最後 Browser 上呈現的就是英文版的 wikipedia Node.js 頁面

<br>

## URI vs. URL vs. URN

* An Uniform Resource Identfier(URI) identifies a resource on the Internet
* An Uniform Resource Locator(URL) is a specialized URI thad identifies a resource `by reachable location`
  * E.g., "http://...", "https://...", "ftp://..."
  * Case sensitive
  * Must be `URL encoded`
    * `"http://host/hello world!"` **→** `"http://host/hello%20world%21"`
* An Uniform Resource Name(URN) is a specialized URI that identifies a resource `by name`
  * E.g., "urn:isbn:0451450523" 

`URI` 沒有明確定義內部的格式，`URL` 必須定義 `protocol Hostname path`，`URN` 則是以名子做為 ID 而沒有明確的位置

<br>

# Web Development

* To build a website that servers web pages
* Code to be run at client side(front end):
  * Dispalys pages
  * Interats with user
  * Sends additional requests
* Code to be run at server side(back end):
  * Handles requests (from multiple client)
  * Stores data
  * Sends "personalized" responses

所謂的(front & back end)最簡單的分辨是你的 Code 是寫在 client 上還是 server 上。

`Front end` 因此便需要處理顯示，與 User 的互動，與 requests

`Back end` 處理所有的 client requests ， 或處理不同的 responses 如行動版、桌面板不同的版本。

<br>

## What'Inside a Web Page?

* Hyper-Text Markup Language(**HTML**):
  * Content and structure
* Cascade Style Sheet (**CSS**):
  * Style (e.g., color, font, width, height, etc.)
* `Javascript`
  * Interactions (e.g., button, textbox, etc.)
  * A general-purpose programming language

<br>

## HTML-Example-index

**`HTML`** ：

```HTML
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <link rel = "stylesheet" href="styles/main.css">
    <title>Document</title>
</head>
<body>
    <h1>This is a Heading</h1>
    <p>This is a paragraph.</p>
    <button id = btn>Show color</button>
    <script type = "text/javascript" src = "scripts/main.js"></script>
</body>
</html>
```

`<!DOCTYPE html>` 定義了以下的 Document，讓 Browser 知道以下的內容是什麼格式。

`<html lang = "en">` 到 `<html>` 這兩個 tag 之間有兩個 child tag `<head>`，`<body>`

`<head>` 不會呈現在 Browser 中， 其中是決定 `<body>` 中的內容怎麼呈現得更好，除了 `<title>`

`<h1>...</h1>` `<p>...</p>` 是其中的內容，與結構位置。在 `HTML` 中只決定結構與內容。

`<script type = "text/javascript" src = "scripts/main.js"></script>` 最後 Load `JS` 內容。

<br>

**`CSS`** ：

```CSS
h1 {
    color: gray;
    text-align: center;
}
```
而 `<h1>` 的顏色，位置，字形則在 `CSS` 中定義。

<br>

**`Javascript`**：

```Javascript
window.onload = function() {
    var btn = document.getElementById("btn");
    btn.addEventListener("click", function(){
        document.body.style.background = randColor();
    });
};
```

取得 `btn` 這個 document 中的 ElementId 放在 `var btn`， `btn.addEventListener` 代表所關注的事件 <br>
`"click"`,發生時執行 `function()` 中的內容，`document.body.style.background = randColor();`<br>
更改了 `body`，`style`，`background` 因此完成按下按鈕後的背景換色。

<br>

最後是 HTML `<meta>` 欄位中的所影響的是編碼模式，IE支援，寬度，viewport的縮放。

![](https://i.imgur.com/UGIHsRf.jpg)

以上範例說明了 `HTML`，`CSS`，`Javascript` 如何影響了 Broswer 如何呈現。

<br>

# HTML

* HTML
  * `<head>` and `<body>`
  * Block vs. lnline elements
  * Lists
  * Links and Attributes
  * Tables
  * Forms

[MDN HTML/Element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element)

## Block vs. lnline elements

```HTML
<div>
    <p>
        Lorem, <b>ipsum</b> <strong>sit</strong> amet consectetur adipisicing 
        elit. <span> Culpa ipsam pariatur distinctio. Corporis ex expedita 
        odit, delectus sint sequi explicabo dolores possimus maxime 
        perspiciatis, officiis omnis! Quam dolorum at minus! </span>
    </p>
</div>
```

`Block element` 即無論如何縮放都會占滿畫面，`Inline element` 可以想像成 `Block` 的單字 <br>
最常見的 `Block element` 基本上就是 Heading，而 Button 就是一種 `Inline element` <br>
在 HTML 中盡量不要使用直接改變 Visual 的方式，這部分應該交給 CSS 負責。

<br>

##  [Lists](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul)

<br>

## Links and Attributes 

```HTML
<link rel = "stylesheet" href="styles/main.css">
```

從 Hello-Html 的 head 中的這段其實就是一種 Attributes，
`rel = "stylesheet" href="styles/main.css"`，設定了這個 Element 的 Attributes

[Attributes](https://developer.mozilla.org/zh-TW/docs/Web/HTML/Attributes) 詳細的 Attributes 查詢

簡單說明最重要的兩個 Element，Link 和 img。

```HTML
<a href = "Page2.html">Go to Page2.</a>
<a href = "Html/Page2.html">Go to Page2.</a>
<img src = "https://i.imgur.com/UGIHsRf.jpg">
```

Link 的使用可以使用絕對位置或相對位置。

<br>

## Tables

```HTML
    <table>
        <tr>
            <th>Headre 1</th>
            <th>Headre 2</th>
        </tr>
        <tr>
            <td>Item 1</td>
            <td>Item 2</td>
        </tr>
        <tr>
            <td>Item 1</td>
            <td>Item 2</td>
        </tr>
    </table>
```

一個最簡單的 Table 範例，在 CSS 的部分講解 Style。

<br>

## Forms & Form Validation


* Queries
  * E.g., specialized presentation for current user
* Passing arguments
  * E.g, user login

<br>

![](https://i.imgur.com/t64TMs9.jpg)

```HTML
<form action="http://sign-in-url" method="get">
    <label for="acc">Username:</label>
    <input id="acc" name = "account" placeholder="Email Address" type="email" required>
    <label for="pw">Password:</label>
    <input id="pw" name = "password" placeholder="8 Digits" type="password" required>
    <button>Login</button>
</form>
```

`http://sign-in-url/?account=Username&password=3345678`

我可以發現，如果使用 GET 在這裡所有的資訊全部在 URL 中顯示，實際上真正要做 Login 的動作 <br>
比較適合使用 POST，實際上 Requests 中除了 Header 和 start 外還有 body 的部分。 <br>

但是今天即便放在 body 中，若是有人進行監聽，依然有機會從封包中捕捉到 Password <br>
所以這部分一般會使用 `https://` 進行加密。

而在 Form 中有 Attributes 可以設定對於 Form Validation，當然這些都只是 HTML 的簡單方式，<br>
在 Js 中可以做到更完整的方式。

<br>

## Assigned Reading and Reference

<br>

[HTML Tutorial](https://www.w3schools.com/html/) W3school 有一些 HTML 的測驗，這裡試著練習一下。

[MDN](https://developer.mozilla.org/zh-TW/docs/Web/HTML) Reference。

<br>

## Exercise

<br>

最後試著完全用 HTML 完成下面一個小的測驗

![](https://i.imgur.com/cD0flx2.jpg)