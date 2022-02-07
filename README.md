# HTML-Learning

**學習關於HTML/CSS/JavaScript筆記**

**參考清大Web-Programming開放式課程**

* How does internet Work?
* Web Development
* HTML
  * "head" and "body"
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

# HTTP Protocol

* A `protocol` is language spoken by machines
  * Defines structure of messages to be exchanged
* HyperText Transfer Protocol(HTTP) defines:
  * Messages : HTTP `request` and HTPP `response`
  * Requests : accessing `resources` (web pages) via `GET POST PUT DELETE` methods
  * Responses : **2**00 OK, **3**01 Moved, **4**04 Not Found, **5**00 Server Error,etc.

關於 [Requests](https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Methods) 與 [Responses](https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status) 詳細的說明

<br>

# HTTP Messages

![](https://i.imgur.com/hyVrpll.png)

* Each HTTP message have
  * `start line`, `header line`, and optionally `body`
* A resource can have different `presentations`
  * HTML, XML, JSON, etc.

從上面的 Example 可以看到 `start` 中 Requests GET 與 Protocol 的版本名稱，Server Respones 200 OK

Requests `Header`，包含自己的 `Host`，`Accept`，`Accept-Encoding` 等代表著接受的端口、檔案格式等資訊

Response `Header`，則有 `Content-Type`，`Content-Length` 代表 Body 的格式，長度等資訊

Body 則是交換的 Resources, 這些交換的 Resources 並不限定於網頁。

<br>

# HTTP Example

![](https://i.imgur.com/E19kX8K.jpg)

從 `Chrome inspect` 上我們就可以看到 HTTP 交換的 Header 與 Body，在這裡從 www.wikipedia.org 送出之後 Response 301 因此 Browser 就轉向尋找 Location，然後就轉向 en.wikipedia.org

# This is an <h1> tag

## This is an <h2> tag

###### This is an <h6> tag

* Item 1
* Item 2
  * Item 2a
  * Item 2b

- [x] This is a complete item
- [ ] This is an incomplete item

`Format one word or one line`

code (4 spaces indent)


# Outline <h1>


