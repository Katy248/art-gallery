document.addEventListener("DOMContentLoaded", function() {
    const loginForm = document.getElementById("login-form");

    loginForm.addEventListener("submit", function(event) {
        event.preventDefault(); // Предотвращаем отправку формы

        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        // Простейшая проверка
        if (username === "" || password === "") {
            alert("Пожалуйста, заполните все поля.");
        } else {
            // Здесь можно добавить код для проверки данных на сервере
            console.log("Логин:", username);
            console.log("Пароль:", password);
            alert("Вы успешно вошли в аккаунт!"); // Успешный вход
            // Здесь можно перенаправить пользователя на другую страницу
            // window.location.href = "dashboard.html"; // Пример перенаправления
        }
    });
});
