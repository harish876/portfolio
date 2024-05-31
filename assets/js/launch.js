document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('form');
    const message = document.getElementById('message');

    form.addEventListener('submit', handleSubmit, false);
    message.addEventListener('keypress', handleEnter, false);
}, false);