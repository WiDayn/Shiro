var logoutLink = document.getElementById('logoutLink');

logoutLink.addEventListener('click', function(e) {
    // 阻止链接的默认行为
    e.preventDefault();

    // 清空token cookie
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';

    // 重定向到登录页面
    window.location.href = './login';
});