document.getElementById('add-image-button').onclick = function() {
    document.getElementById('image-modal').style.display = 'block';
};

document.getElementsByClassName('close')[0].onclick = function() {
    document.getElementById('image-modal').style.display = 'none';
};

document.getElementById('submit-image').onclick = function() {
    const imageUrl = document.getElementById('image-url').value;
    const imageDescription = document.getElementById('image-description').value;

    if (imageUrl) {
        const gallery = document.getElementById('gallery');

        const card = document.createElement('div');
        card.className = 'card';
        card.innerHTML = `
            <img src="${imageUrl}" alt="Image">
            <p>${imageDescription}</p>
        `;
        
        gallery.appendChild(card);
        document.getElementById('image-url').value = '';
        document.getElementById('image-description').value = '';
        document.getElementById('image-modal').style.display = 'none';
    } else {
        alert('Пожалуйста, введите URL изображения.');
    }
};

// Закрытие модального окна при клике вне его
window.onclick = function(event) {
    const modal = document.getElementById('image-modal');
    if (event.target === modal) {
        modal.style.display = 'none';
    }
};

