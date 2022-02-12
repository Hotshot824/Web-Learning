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

<br>

## [Google font](https://fonts.google.com/)

Google提供了許多免費字體供使用，在Page中使用link就可以在CSS中使用字體。

當然增加字體是會影響Loadtime，盡量在字體使用上使用Less is more的原則。

<br>

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

<br>

## Chrome Inspector

![](https://i.imgur.com/JOlCDI1.jpg)

在Inspector中能看到所有影響這個element的CSS Rule，並且可以在Inspector中能動態的開關這些Rule。
並不會影響到source code。

<br>

## More Selector

```CSS
/* composition*/
#todo-list, li.done {...}
/* descendant selector*/
li a {...}
/* adjacent selector*/
li.done + li {...}
/* attribute selector*/
a[href="http://..."] {...}
```
![](https://i.imgur.com/TeKrgNW.jpg)


在CSS中選擇影響風格的對象，也可以指向特定attribute的element。

<br>

## Pseudo Classes & Elements

```CSS
/* pseudo class selector */
a:hover, a:visited {...}
ul:nth-of-type(3) {...}
/* pseudo element selector */
p::first-latter {
  font-size: 200%;
}
h1::before{
  content: url(image.gif);
}
```

* More [selector examples](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors)

<br>

## Inheritance

* Most style properties of an element will be **inherited** by its descendants
  * E.g., font, color, text-align, etc.
* Common excpetions are those related to the box model
  * E.g., height, width, border-width, etc.
* Check [this refrence](https://developer.mozilla.org/en-US/docs/Web/CSS/inherit) to see if a property is inheritable.

最直覺判斷是否能被繼承的判斷，如顏色，字形這種單純風格樣式的類型可被繼承，而寬高 <br>
這些與Layout有關的則不能被繼承。

## Cascading

* Final properties gotten by an element are **cascaded** from all applicable rules.

```HTML
<body>
  <ul id="todo-list">
      <li>...</li>
      <li class="done">
        TODO
      </li>
    </ul>
 </body>
```

```CSS
body{
    background-color: black;
    color: white;
  }
  #todo-list{
    font-weight: bold;
  }
  #todo-list li{
    color: red;
  }
  li.done{
    color: blue;
    text-decoration: line-throuhg;
  }
```

![](https://i.imgur.com/5Og1eDZ.jpg)

從上面的例子中其實所有的Rule都能繼承在TODO上，但我們能看到最終結果上 <br>
`<body>`，`<#todo-list>`，`li.done`，最終是`<#todo-list>`呈現在畫面上，這就是cascading。 <br>
最後的`color: blue;`實際上並未發生。

<br>

## How to Resovle Conflicts?

基本上在CSS中有三種方式計算誰是最後呈現的結果。

* By importance
```CSS
body{ 
  color: gray !important; 
}
```
加入 !important 必然使用。
* By **specificity** 

![](https://i.imgur.com/HhIgC6Q.jpg)

從左往右計數，`ID Selectors`是最大的因數，只要有`ID Selectors`就必然使用 <br>
然後是`Class`最後是`Type`，所以是 `ID` > `Class` > `Type`。
* By source order
  * Rule written later win

最後如果同分或者是完全相同，就以編寫時的順序最後的決定。

<br>

# Layout

## HTML Rendering
* The content are rendered following the **normal flow**
  * Block element are laid out vertically
  * Inline element are laid out horizontally





