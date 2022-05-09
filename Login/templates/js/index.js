document.querySelector("#checkbtn").addEventListener("click", function () {
    console.log("function called check button");

    fetch("ping", {
        method: "POST",
    })
        .then((response) => {
            return response.json();
        })
        .then((response) => {
            alert(typeof(response))
            console.log(response);
        })
        .catch((error) => {
            console.log(`Error: ${error}`);
        })
}, false);