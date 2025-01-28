document.addEventListener("DOMContentLoaded", () => {
    const toggleButton = document.getElementById("toggle-sidebar");
    const dropdownMenu = document.getElementById("dropdown-menu");

    const activeClass = "active";

    const showMenu = (menu) => {
        menu.classList.add(activeClass);
    }

    const hideMenu = (menu) => {
        menu.classList.remove(activeClass);
    }
    // Обработчик клика
    toggleButton.addEventListener("click", () => {
        if (!dropdownMenu.classList.contains(activeClass)) {
            showMenu(dropdownMenu)
        } else {
            hideMenu(dropdownMenu)
        }
    });

    // Закрытие меню при клике вне
    document.addEventListener("click", (event) => {
        if (!toggleButton.contains(event.target) && !dropdownMenu.contains(event.target)) {
            hideMenu(dropdownMenu)
        }
    });
});
