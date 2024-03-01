document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault(); // 阻止表单的默认提交行为

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/user/login', true);
    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(this.responseText);
            if (response.message === "Login successful") {
                // 保存token到cookie
                document.cookie = "token=" + response.token + ";path=/";
                document.location.href = './dashboard';
            }
        }
    };

    var data = JSON.stringify({
        "username": username,
        "password": password
    });

    xhr.send(data);
});