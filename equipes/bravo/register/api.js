const API_BASE = 'http://localhost:8080/api/v1';

document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('form');
    if (form) {
        form.addEventListener('submit', registrar);
    }
});

async function registrar(event) {
    event.preventDefault();

    const username = document.getElementById('username').value.trim();
    const email = document.getElementById('email').value.trim();
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirm-password').value;

    if (password !== confirmPassword) {
        alert('As senhas não coincidem.');
        return;
    }

    try {
        const response = await fetch(`${API_BASE}/users`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: username, email, password })
        });

        if (!response.ok) {
            const erro = await response.json();
            alert(erro.error || 'Erro ao cadastrar. Tente novamente.');
            return;
        }

        alert('Conta criada com sucesso!');
        window.location.href = '../index.html';
    } catch (error) {
        console.error('Erro ao registrar:', error);
        alert('Não foi possível conectar ao servidor.');
    }
}
