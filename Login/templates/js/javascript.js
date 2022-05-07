window.onload = function () {
    console.log("function called options add");

    var obj = document.getElementById("inputState");
    obj.options.add(new Option("Taiwan", "taiwan"))
    obj.options.add(new Option("China", "china"))
    obj.options.add(new Option("Thailand", "thailand"))
    obj.options.add(new Option("Korean", "korean"))
    obj.options.add(new Option("HongKong", "hongkong"))
    obj.options.add(new Option("AAAA", "AAAA"))
}

document.querySelector("#signup").addEventListener("submit", function (e) {
    console.log("function called send formData");
    e.preventDefault()

    alert("Data has been send");
    let form = new FormData(document.querySelector('form'));
    fetch("sign_up", {
        method: "POST",
        body: form
    });
    window.location.reload();
}, false);

document.querySelector("#signin").addEventListener("submit", function (e) {
    console.log("function called send formData");
    e.preventDefault()

    alert("Data has been send");
    let form = new FormData(document.querySelector('form'));
    fetch("sign_in", {
        method: "POST",
        body: form
    });
    window.location.reload();
}, false);