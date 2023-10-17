package static

var CreateTxForm = `
<!DOCTYPE html>
<html lang="ru">
<meta charset="UTF-8" />
<head>
    <title>Форма для перевода</title>
</head>
<body>
<h1>Перевод транзакции</h1>
<form id="transactionForm">
    <label for="sender">От кого:</label>
    <input type="text" id="sender" name="sender" required><br><br>

    <label for="receiver">Кому:</label>
    <input type="text" id="receiver" name="receiver" required><br><br>

    <label for="amount">Сколько:</label>
    <input type="number" id="amount" name="amount" required><br><br>

    <input type="submit" value="Перевести">
</form>

<script>
    // Функция для отправки данных формы на бэкенд
    function sendTransaction(event) {
        event.preventDefault(); // Предотвращаем стандартное поведение отправки формы

        const form = document.getElementById('transactionForm');
        const formData = new FormData(form);

        fetch('http://localhost:8000/process-transaction', {
			headers: {
				  'Accept': 'application/json',
				  'Content-Type': 'application/json'
				},
            method: 'POST',
            body: JSON.stringify(Object.fromEntries(formData))
        })
            .then(response => response.json())
            .then(data => {
                // Обработка ответа от сервера
                console.log(data);
                // Можно добавить дополнительную обработку, например, отображение сообщения об успешной транзакции
            })
            .catch(error => {
                // Обработка ошибок
                console.error('Произошла ошибка:', error);
            });
    }

    const transactionForm = document.getElementById('transactionForm');
    transactionForm.addEventListener('submit', sendTransaction);
</script>
</body>
</html>
`
