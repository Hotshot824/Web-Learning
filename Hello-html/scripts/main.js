window.onload = function() {
    var btn = document.getElementById("btn");
    btn.addEventListener("click", function(){
        document.body.style.background = randColor();
    });
};

function randColor() {
    function co(lor) {
        const digits = [0,1,2,3,4,5,6,7,8,9,'a','b','c','d','e','f'];
        lor += digits[ Math.floor(Math.random()*16) ];
        if (lor.length === 6) {
            return lor;
        } else {
            return co(lor);
        }
    }
    return '#' + co('');
}