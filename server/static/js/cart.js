document.addEventListener('DOMContentLoaded', function() {
    function showNotification(message, type = 'success') {
        const notification = document.createElement('div');
        notification.className = `notification ${type}`;
        notification.textContent = message;
        document.body.appendChild(notification);

        setTimeout(() => {
            notification.remove();
        }, 3000);
    }

    const deleteButtons = document.querySelectorAll('.btn-delete');
    deleteButtons.forEach(button => {
        button.addEventListener('click', async function() {
            const cartItem = this.closest('.cart-item');
            const productId = cartItem.dataset.productId;

            try {
                const response = await fetch(`/api/cart/${productId}`, {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });

                if (response.ok) {
                    cartItem.remove();
                    showNotification('Товар удален из корзины');
                    updateCartSummary();
                } else {
                    throw new Error('Ошибка при удалении товара');
                }
            } catch (error) {
                showNotification(error.message, 'error');
            }
        });
    });

    const quantityInputs = document.querySelectorAll('.cart-item-quantity input');
    quantityInputs.forEach(input => {
        input.addEventListener('change', async function() {
            const cartItem = this.closest('.cart-item');
            const productId = cartItem.dataset.productId;
            const quantity = parseInt(this.value);

            if (quantity < 1) {
                this.value = 1;
                return;
            }

            try {
                const response = await fetch(`/api/cart/${productId}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ quantity })
                });

                if (response.ok) {
                    updateCartSummary();
                    showNotification('Количество товара обновлено');
                } else {
                    throw new Error('Ошибка при обновлении количества');
                }
            } catch (error) {
                showNotification(error.message, 'error');
            }
        });
    });

    async function updateCartSummary() {
        try {
            const response = await fetch('/api/cart/summary');
            if (response.ok) {
                const data = await response.json();
                document.querySelector('.total-amount').textContent = `${data.total} ₽`;
                document.querySelector('.total-items').textContent = data.itemsCount;
            }
        } catch (error) {
            console.error('Ошибка при обновлении итоговой суммы:', error);
        }
    }

    const checkoutButton = document.querySelector('.btn-primary');
    if (checkoutButton) {
        checkoutButton.addEventListener('click', async function() {
            try {
                const response = await fetch('/api/orders', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });

                if (response.ok) {
                    showNotification('Заказ успешно оформлен');
                    setTimeout(() => {
                        window.location.href = '/orders';
                    }, 1500);
                } else {
                    throw new Error('Ошибка при оформлении заказа');
                }
            } catch (error) {
                showNotification(error.message, 'error');
            }
        });
    }

    updateCartSummary();
}); 