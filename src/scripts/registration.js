document.addEventListener("DOMContentLoaded", function() {
    const registrationForm = document.getElementById("registration-form");

    registrationForm.addEventListener("submit", function(event) {
        event.preventDefault(); // Предотвращаем отправку формы

        const username = document.getElementById("username").value;
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;
        const confirmPassword = document.getElementById("confirm-password").value;

        // Проверка заполненности полей
        if (username === "" || email === "" || password === "" || confirmPassword === "") {
            alert("Пожалуйста, заполните все поля.");
            return;
        }

        // Проверка совпадения паролей
        if (password !== confirmPassword) {
            alert("Пароли не совпадают. Пожалуйста, убедитесь, что вы ввели одинаковые пароли.");
            return;
        }

        // Здесь можно добавить код для отправки данных на сервер
        console.log("Логин:", username);
        console.log("Почта:", email);
        console.log("Пароль:", password);
        alert("Вы успешно зарегистрированы!"); // Успешная регистрация

        // Здесь можно перенаправить пользователя на другую страницу
        // window.location.href = "login.html"; // Пример перенаправления
    });
});
