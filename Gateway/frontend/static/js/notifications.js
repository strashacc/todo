window.onload = function () {
    fetch('/api/notifications')
        .then(res => res.json())
        .then(data => {
            const ul = document.getElementById('notification-list');
            if (ul && data.length > 0) {
                data.forEach(n => {
                    const li = document.createElement('li');
                    li.textContent = n.text;
                    ul.appendChild(li);
                });
            } else if (ul) {
                ul.innerHTML = '<li>Нет уведомлений</li>';
            }
        })
        .catch(() => {
            const ul = document.getElementById('notification-list');
            if (ul) ul.innerHTML = '<li>Ошибка при загрузке уведомлений</li>';
        });
};
