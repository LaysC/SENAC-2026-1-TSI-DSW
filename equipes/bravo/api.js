const API_BASE = 'http://localhost:8080/api/v1';

document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('form');
    if (form) {
        form.addEventListener('submit', login);
    }
});

async function login(event) {
    event.preventDefault();

    const username = document.getElementById('username').value.trim();
    const password = document.getElementById('password').value;

    if (!username || !password) return;

    try {
        const response = await fetch(`${API_BASE}/users/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name: username, password })
        });

        if (!response.ok) {
            const erro = await response.json();
            alert(erro.error || 'Usuário ou senha inválidos.');
            return;
        }

        const data = await response.json();
        localStorage.setItem('user_id', data.user_id);
        localStorage.setItem('username', data.name);

        window.location.href = './todolist/index.html';
    } catch (error) {
        console.error('Erro ao fazer login:', error);
        alert('Não foi possível conectar ao servidor.');
    }
}
