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