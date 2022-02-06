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

<br>

* `https://` : **Protocol** 使用何種通訊協定

* `en.wikipedia.org` : **Domain name** 連線至哪一個網域

* `/wiki/Main_Page` : **Path** 網域所能提供的 Resources

<br>

首先將Domain Name 拿出詢問 DNS 這個Domain Name所代表的 IP`(Internet Protocol)`

`en.wikipedia.org` 是給 User 使用，如`www.wikipedia`代表不同語言的服務。

而 Machine 所使用的則是IP Address。 之後 Browser 才可以連到對應的 server 。

這時便可以與 Server 建立 Connection 可能是 UDP 或 TCP。  

<br>
 
![](https://i.imgur.com/2ZXNkNT.png)

當與 Server 連線後便可以使用 HTTP 使用 `request` 如 GET/documentart.mpe

Server 回復 200 並將 Page 回傳，這樣就可以在 Browser 上呈現。

這就是按下 Enter 之後到 Browser 上呈現背後所經過的流程。

<br>

# HTTP

* A `protocol` is language spoken by machines
  * Defines structure of messages to be exchanged
* HyperText Transfer Protocol(HTTP) defines:
  * Messages : HTTP `request` and HTPP `response`
  * Requests : accessing resources (web pages) via `GET POST PUT DELETE` methods







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


