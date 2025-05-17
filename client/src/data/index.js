const CONTACTS = [{ name: "Почта автора", email: "petrovanton247@gmail.com", url: "mailto:petrovanton247@gmail.com" }];
const LINKS = [
    { name: "Исходные коды", url: "https://gitlab.com/side-projects133505/art-gallery", icon: "fab fa-github" },
    { name: "Профиль администратора", url: "http://80.74.25.113:5173/user/1" },
];
const CHANGELOGS = [
    {
        title: "Версия 0.0.4",
        preview: true,
        changes: ["Обновлено руководство", "Обновлены ссылки", "Добавлен список пользователей в панели администратора", "Обновлены стили кнопок (в основном добавлены transition-ы)"],
    },
    { title: "Версия 0.0.3", changes: ["Добавлено описание пользователя", "Добавлено изменение текста публикации", "Исправлена проблема при создании публикации"] },
    { title: "Версия 0.0.2", changes: ["Добавлено удаление постов", "Добавлена роль администратора", "Добавлено отображение времени публикации (не только даты)"] },
];

export { CONTACTS, LINKS, CHANGELOGS };
