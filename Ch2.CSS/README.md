# CSS

## [CSS Zen Garden](http://www.csszengarden.com/)

一個內容相同的HTML，不同的CSS可以做出完全不同的網頁。

<br>

## Outline

* The Basics
* Selectors
* Layout
* Stacking Order

<br>

# The Basics

## Grammar

```CSS
/*for all h1's & img's*/
h1 {
    color:red;
    font-size: 56px;
}
img {
    border-color: blue;
    border-width: 3px;
}
```

## Color and Background

```CSS
h1{
    color: rgb(255, 0, 0);
    font-size: 56px;
    background-color: rgb(255, 255, 255);
}
body{
    background-image: url(../images/bg.jpg);
    background-size: cover;
}
```

![](https://i.imgur.com/4iuCegb.jpg)


從CSS基本給color和background的方式。

<br>

## CSS Properties

* Background:
  * E.g., background-color, background-image, etc.
* Text:
  * E.g., color, text-align, text-decoration, etc.
* Background:
  * E.g., font-family, font-size, font-style, font-weight, etc.
* Background:
  * E.g., list-style-type, list-style-image, vertical-align, etc.
* Background:
  * E.g., width, height, border, padding, margin, display, <br>
  visbility, float, position, etc.
* See [a list here](https://www.tutorialrepublic.com/css-reference/css3-properties.php)

<br>

## Native Font Stack

```CSS
body{
  font-family:
  /*Safari for OS X and iOS */
  -apple-system,
  /*Chrome >= 56 for OS X, Windows, Linux and Android */
  system-ui,
  /*Chrome < 56 for OX X */
  BlinkMacSystemFont,
  /*Windows*/
  "Segoe UI",
  /*Android*/
  "Roboto"
  /*Basic web fallback */
  "Helvetica Neue", Arial, sans-seruf;
}
```

![](https://i.imgur.com/a8N5R3b.jpg)

## [Google font](https://fonts.google.com/)

Google提供了許多免費字體供使用，在Page中使用link就可以在CSS中使用字體。

當然增加字體是會影響Loadtime，盡量在字體使用上使用Less is more的原則。

# Selectors

## ID & element & class selector

```HTML
<ul id="todo-list">
    <li class="done">TODO 1</li>
    <li>TODO 2</li>
    <li class="done">TODO 3</li>
</ul>
```
```CSS
li{                                   /* element/tag selector*/
    font-weight: bold;
}
#todo-list{                           /* ID selector*/
    background-color: gray;
}
.done {                               /* class selector*/
    text-decoration: line-through;
}
```

![](https://i.imgur.com/xoiEyvb.jpg)

在CSS中如何控制selector有三種方是，上面是範例。

## Chrome Inspector

![](https://i.imgur.com/JOlCDI1.jpg)

在Inspector中能看到所有影響這個element的CSS Rule，並且可以在Inspector中能動態的開關這些Rule。
並不會影響到source code。