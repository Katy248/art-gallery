// JavaScript для выпадающего меню

document.addEventListener("DOMContentLoaded", () => {
    const uploadForm = document.getElementById("upload-form");
    const fileUpload = document.getElementById("file-upload");
    const descriptionInput = document.getElementById("description");
    const imageGallery = document.getElementById("image-gallery");

    const modal = document.getElementById("modal");
    const modalContent = document.getElementById("modal-media");
    const closeModal = document.getElementById("close-modal");

    // Функция открытия модального окна
    function openModal(mediaElement) {
        modalContent.innerHTML = ""; // Очистить содержимое
        modalContent.appendChild(mediaElement);
        modal.classList.remove("hidden");
    }

    // Закрытие модального окна
    closeModal.addEventListener("click", () => {
        modal.classList.add("hidden");
        modalContent.innerHTML = "";
    });

    // Закрытие по клику на фон
    modal.addEventListener("click", (e) => {
        if (e.target === modal) {
            modal.classList.add("hidden");
            modalContent.innerHTML = "";
        }
    });

    uploadForm.addEventListener("submit", (event) => {
        event.preventDefault();

        const file = fileUpload.files[0];
        const description = descriptionInput.value;

        if (!file) {
            alert("Выберите файл для загрузки!");
            return;
        }

        const fileURL = URL.createObjectURL(file);

        const galleryItem = document.createElement("div");
        galleryItem.classList.add("gallery-item");

        let mediaElement;

        if (file.type.startsWith("image/")) {
            mediaElement = document.createElement("img");
            mediaElement.src = fileURL;
            mediaElement.alt = description;
            mediaElement.style.width = "100%";
            mediaElement.style.borderRadius = "10px";
        } else if (file.type.startsWith("video/")) {
            mediaElement = document.createElement("video");
            mediaElement.src = fileURL;
            mediaElement.controls = true;
            mediaElement.style.width = "100%";
            mediaElement.style.borderRadius = "10px";
        }

        // Добавить клик для открытия модального окна
        galleryItem.addEventListener("click", () => {
            const clonedMedia = mediaElement.cloneNode(true);
            clonedMedia.style.width = "100%"; // Увеличиваем размер в модалке
            clonedMedia.style.height = "auto";
            clonedMedia.controls = true;
            openModal(clonedMedia);
        });

        galleryItem.appendChild(mediaElement);

        if (description) {
            const descText = document.createElement("p");
            descText.textContent = description;
            descText.style.marginTop = "10px";
            descText.style.fontSize = "14px";
            descText.style.color = "#555";
            galleryItem.appendChild(descText);
        }

        imageGallery.appendChild(galleryItem);

        uploadForm.reset();
    });
});

