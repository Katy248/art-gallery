document.getElementById('toggle-sidebar').addEventListener('click', function() {
    const sidebar = document.getElementById('sidebar');
    sidebar.style.display = sidebar.style.display === 'block' ? 'none' : 'block';
});
