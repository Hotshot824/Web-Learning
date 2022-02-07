# HTML-Learning

**學習關於HTML/CSS/JavaScript筆記**

**參考清大Web-Programming開放式課程**

* How does internet Work?
* Web Development
* HTML
  * head and body
  * Block vs. lnline elements
  * Lists
  * Links and Attributes
  * Tables
  * Forms

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

# Hello-Html Example

以下是一個完整的 `HTML` 格式：

```HTML
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
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

<br>

以下是一個完整的 `CSS` 範例：

```CSS
h1 {
    color: gray;
    text-align: center;
}
```
而 `<h1>` 的顏色，位置，字形則在 `CSS` 中定義。

<br>