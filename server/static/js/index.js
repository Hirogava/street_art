const successMessage = document.getElementById('success-message');

async function addToCart(button) {
    try {
        const productId = button.getAttribute('data-product');

        const response = await fetch('/api/product/cart', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id: parseInt(productId, 10),
                count: 1
            })
        });

        const result = await response.json();

        if (result.status === "ok") {
            successMessage.style.display = "block";
            
            setTimeout(() => {
                successMessage.style.animation = "fadeOut 0.3s ease-out";
                setTimeout(() => {
                    successMessage.style.display = "none";
                    successMessage.style.animation = "slideInRight 0.3s ease-out";
                }, 300);
            }, 3000);
        } else {
            alert('Ошибка добавления в корзину');
        }
    } catch (err) {
        console.error('Ошибка:', err);
        alert('Ошибка при отправке запроса');
    }
}