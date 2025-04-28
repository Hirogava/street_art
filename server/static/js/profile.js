function openModal() {
    console.log('Opening modal...');
    const depositModal = document.getElementById('depositModal');
    if (!depositModal) {
        console.error('Modal element not found!');
        return;
    }
    depositModal.style.display = 'flex';
    document.body.style.overflow = 'hidden';
    console.log('Modal opened');
}

function closeModal() {
    console.log('Closing modal...');
    const depositModal = document.getElementById('depositModal');
    if (!depositModal) {
        console.error('Modal element not found!');
        return;
    }
    
    depositModal.style.opacity = '0';
    const modalContent = depositModal.querySelector('.modal-content');
    if (modalContent) {
        modalContent.style.transform = 'translateY(-20px)';
    }
    
    setTimeout(() => {
        depositModal.style.display = 'none';
        document.body.style.overflow = 'auto';
        depositModal.style.opacity = '1';
        if (modalContent) {
            modalContent.style.transform = 'translateY(0)';
        }
    }, 300);
    console.log('Modal closed');
}

function showNotification(message, type = 'error') {
    const notification = document.getElementById('notification');
    notification.textContent = message;
    notification.className = `notification ${type}`;
    notification.classList.add('active');
    
    setTimeout(() => {
        notification.classList.remove('active');
    }, 3000);
}

async function deposit() {
    const amountInput = document.getElementById('amount');
    const amount = parseFloat(amountInput.value);

    if (!amount || amount <= 0) {
        showNotification('Пожалуйста, введите корректную сумму');
        return;
    }

    try {
        const formData = new FormData();
        formData.append('money', amount.toString());

        const response = await fetch('/api/balance/deposit', {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            const error = await response.json();
            showNotification(error.message || 'Ошибка пополнения баланса');
            return;
        }

        const result = await response.json();
        showNotification('Баланс успешно пополнен!', 'success');
        
        const balanceElement = document.querySelector('.balance-amount');
        if (balanceElement) {
            balanceElement.textContent = result.new_balance.toFixed(2) + ' ₽';
        }
        
        closeModal();
        amountInput.value = '';
    } catch (error) {
        console.error('Ошибка:', error);
        showNotification('Ошибка соединения с сервером');
    }
}

document.addEventListener('DOMContentLoaded', function() {
    const depositModal = document.getElementById('depositModal');

    depositModal.addEventListener('click', function(event) {
        if (event.target === depositModal) {
            closeModal();
        }
    });

    document.addEventListener('keydown', function(event) {
        if (event.key === 'Escape' && depositModal.style.display === 'flex') {
            closeModal();
        }
    });
}); 