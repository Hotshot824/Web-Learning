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

<br>

# HTTP Messages

* Each HTTP message have
  * `lnital line`, `header line`, and optionally `body`





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


