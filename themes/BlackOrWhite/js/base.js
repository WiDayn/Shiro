function getTokenFromCookie() {
    // 从document.cookie获取所有cookie，并转换为一个对象
    const cookies = document.cookie.split(';').reduce((acc, cookie) => {
        const [key, value] = cookie.split('=').map(c => c.trim());
        acc[key] = value;
        return acc;
    }, {});

    // 返回名为'token'的cookie，如果不存在则返回null
    return cookies.token || null;
}