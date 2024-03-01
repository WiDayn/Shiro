document.getElementById('submitBtn').addEventListener('click', function() {
    event.preventDefault(); // 阻止表单的默认提交行为

    var username = document.getElementById('username').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;

    // 确认用户接受了注册协议
    var termsAccepted = document.getElementById('agreementCheck').checked;
    if (!termsAccepted) {
        alert("您必须接受注册协议才能注册。");
        return;
    }

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/user/register', true);
    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.onload = function() {
        if (this.status === 200) {
            // 注册成功，跳转到登录页面
            window.location.href = "/admin/login";
        } else {
            // 处理错误，显示错误信息等
            alert("注册失败，请稍后再试。");
        }
    };

    var data = JSON.stringify({
        "username": username,
        "password": password,
        "email": email
    });

    xhr.send(data);
});