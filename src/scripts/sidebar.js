document.addEventListener("DOMContentLoaded", () => {
    const toggleButton = document.getElementById("toggle-sidebar");
    const dropdownMenu = document.getElementById("dropdown-menu");

    // Обработчик клика
    toggleButton.addEventListener("click", () => {
        if (dropdownMenu.style.display === "block") {
            dropdownMenu.style.display = "none"; // Скрыть меню
        } else {
            dropdownMenu.style.display = "block"; // Показать меню
        }
    });

    // Закрытие меню при клике вне
    document.addEventListener("click", (event) => {
        if (!toggleButton.contains(event.target) && !dropdownMenu.contains(event.target)) {
            dropdownMenu.style.display = "none"; // Скрыть меню
        }
    });
});
