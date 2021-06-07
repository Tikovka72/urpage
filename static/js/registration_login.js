function sendRegistration() {
    if (!checkPasswordsMatch()) {
        return false
    }

    let status;

    let data = new FormData();

    data.append("username", document.getElementById("username").value);
    data.append("email", document.getElementById("email").value)
    data.append("password", document.getElementById("password").value)

    fetch("/do/registration", {method: 'post', body: data}).then(function (r) {
        status = r.status

        if (status === 200) {
            r.json().then(function (j) {
                if (j["Err"] === "") {
                    window.location.replace("/login")
                } else {
                    setError(j["Err"])
                }
            })
        } else {
            alert("что-то пошло не так...");
        }
    })

    return false
}

function sendLogin() {
    let status;

    let data = new FormData();

    data.append("email", document.getElementById("email").value)
    data.append("password", document.getElementById("password").value)

    fetch("/do/login", {method: 'post', body: data}).then(function (r) {
        status = r.status

        if (status === 200) {
            console.log(r)
            r.json().then(function (j) {
                console.log(j)
                if (j["Err"] === "") {
                    window.location.replace("/")
                } else {
                    setError(j["Err"])
                }
            })
        } else {
            alert("что-то пошло не так...");
        }
    })

    return false
}