document.addEventListener('DOMContentLoaded', function() {
    const editProfileForm = document.getElementById('editProfileForm');
    const notification = document.getElementById('notification');

    function showNotification(message, type = 'error') {
        notification.textContent = message;
        notification.className = `notification ${type}`;
        notification.classList.add('active');
        
        setTimeout(() => {
            notification.classList.remove('active');
        }, 3000);
    }

    editProfileForm.addEventListener('submit', async function(event) {
        event.preventDefault();

        const formData = new FormData(editProfileForm);
        const submitButton = editProfileForm.querySelector('button[type="submit"]');
        const originalButtonText = submitButton.textContent;

        const newPassword = formData.get('new_password');
        const confirmPassword = formData.get('confirm_password');
        
        if (newPassword && newPassword !== confirmPassword) {
            showNotification('Пароли не совпадают');
            return;
        }

        try {
            submitButton.disabled = true;
            submitButton.innerHTML = '<span class="loading-spinner"></span>';

            const response = await fetch('/api/profile/edit_profile', {
                method: 'POST',
                body: formData
            });

            const result = await response.json();

            if (response.ok) {
                showNotification('Профиль успешно обновлен! Перенаправление...', 'success');
                setTimeout(() => {
                    window.location.href = '/profile';
                }, 1000);
            } else {
                showNotification(result.message || 'Ошибка обновления профиля');
            }
        } catch (error) {
            console.error('Ошибка:', error);
            showNotification('Ошибка соединения с сервером');
        } finally {
            submitButton.disabled = false;
            submitButton.textContent = originalButtonText;
        }
    });
}); 