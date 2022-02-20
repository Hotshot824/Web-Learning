# Javascript: The Basics

## HTML vs CSS vs Javascript

* HTML deinfe content and element structure
  * The "Nouns"
* CSS define how an element looks
  * The "Adjectives"
* Javascript define how an element interact
  * The "Verbs"


# Javascript

* An implementation of **ECMAScript(ES)** standard
  * Javascript 1.5, 1.7, 1.8.5 are non-offficail standards maintained by Mozilla
* ES5 = ECMAScript 5 (2009)
  * Supported by major browsers
  * Covered by this class
* ES6 = ECMAScript 6 = ES2015
* ES7 = ECMAScript 7 = ES2016
  * Not full supported yet
  * Luckily, transpilers such as Babel are avalabie
  * To be covered the next class

## Running Javascript in Browsers

* In *.js file
  * When loading HTML, code inside `<script>` is executed immediately
```javascript
window.onload = function(){
    var el = document.querySelector('h1');
    el.textContent = ('Hello JS!');
}
```

* In Chrome console
  * console.log(el.textContent);
