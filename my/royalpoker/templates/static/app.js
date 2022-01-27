
// 添加登录绑定
loginButton = (event) => {
    let username = $("#username").val();
    let password = $("#password").val();
    let data = {username: username, password: password}
    $.ajax({
        url: 'http://localhost:8080/login',
        type: 'POST',
        data: JSON.stringify(data),
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        async: false,
        success: function (res) {
            if (res.code !== 0) {
                alert(res.msg)
                return
            }

            app.user.id = res.data.id
            app.user.name = res.data.username
            app.user.is_admin = res.data.is_admin
            app.showDiv = 'index'

            localStorage.setItem('token', res.data.token);
        }
    });
}
// 添加退出绑定
logoutButton = (event) => {
    localStorage.removeItem('token')
    app.user = {}
    app.showDiv = 'login'
}