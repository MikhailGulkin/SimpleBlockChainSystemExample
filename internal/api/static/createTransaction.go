package static

const CreateTxForm = `
<!DOCTYPE html>
<html lang="ru">
<meta charset="UTF-8"/>
<head>
    <title>Форма для перевода</title>
</head>
<style>
    .container {
        display: flex;
        flex-direction: column;
        text-align: center;
        padding-left: 100px;
    }

    .item {
    }

</style>
<body>
<div class="container">
    <div class="item">
        <h2>Перевод транзакции</h2>
        <form id="transactionForm">
            <label for="sender">От кого:</label>
            <input type="text" id="sender" name="sender" required><br><br>

            <label for="receiver">Кому:</label>
            <input type="text" id="receiver" name="receiver" required><br><br>

            <label for="amount">Сколько:</label>
            <input type="number" id="amount" name="amount" required><br><br>

            <input type="submit" value="Перевести">
        </form>
        <div>
            <p id="transactionResultField"></p>
        </div>
    </div>

    <div class="item">
        <h2>Узнать статус транзакции</h2>
        <form id="checkStatusTransactionForm">
            <label for="checkTxStatus">Айди транзакции:</label>
            <input type="text" id="checkTxStatus" name="transactionId" required><br><br>

            <input type="submit" value="Узнать статус">
        </form>
        <div>
            <p id="checkStatusTransactionResultField"></p>
        </div>
    </div>

    <div class="item">
        <h2>Обработать все транзакции и создать блок(Майнить)</h2>
        <form id="mineForm">
            <label for="mine">Адресс майнера:</label>
            <input type="text" id="mine" name="address" required><br><br>

            <input type="submit" value="Майнить">
        </form>
        <div>
            <p id="mineResultField"></p>
        </div>
    </div>
    <div class="item">
        <h2>Получить все кошельки</h2>
        <button id="getWallets">Получить</button>
        <div>
            <ul id="wallets"></ul>
        </div>
    </div>

    <div class="item">
        <h2>Пороверить валидность блокчейна</h2>
        <button id="checkBC">Проверить</button>
        <div>
            <p id="checkStatusBlockChainField"></p>
        </div>
    </div>

</div>


<script>
    const transactionResultField = document.getElementById('transactionResultField');
    const checkStatusTransactionResultField = document.getElementById('checkStatusTransactionResultField');
    const mineResultField = document.getElementById('mineResultField');
    const wallets = document.getElementById('wallets');
    const blockChain = document.getElementById('checkStatusBlockChainField');

    function sendTransaction(event) {
        event.preventDefault();
        transactionResultField.innerHTML = '';

        const form = document.getElementById('transactionForm');

        const formData = new FormData(form);
        // form.reset();


        fetch('http://localhost:8000/process-transaction', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: 'POST',
            body: JSON.stringify(Object.fromEntries(formData))
        })
            .then(response => response.text())
            .then(data => {
                transactionResultField.innerHTML = JSON.stringify(data, null, 4);
            })
            .catch(error => {
                console.error('Произошла ошибка:', error);
            });
    }

    function checkStatusTransaction(event) {
        event.preventDefault();
        checkStatusTransactionResultField.innerHTML = '';

        const form = document.getElementById('checkStatusTransactionForm');

        const formData = new FormData(form);
        form.reset();

        fetch('http://localhost:8000/check-tx-status', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: 'POST',
            body: JSON.stringify(Object.fromEntries(formData))
        })
            .then(response => response.json())
            .then(data => {
                checkStatusTransactionResultField.innerHTML = JSON.stringify(data, null, 4);
            })
            .catch(error => {
                console.error('Произошла ошибка:', error);
            });

    }

    function mine(event) {
        event.preventDefault();
        mineResultField.innerHTML = '';

        const form = document.getElementById('mineForm');

        const formData = new FormData(form);
        form.reset();

        fetch('http://localhost:8000/mine', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: 'POST',
            body: JSON.stringify(Object.fromEntries(formData))
        })
            .then(response => response.json())
            .then(data => {
                mineResultField.innerHTML = JSON.stringify(data, null, 4);
            })
            .catch(error => {
                console.error('Произошла ошибка:', error);
            });

    }

    function getWallets(event) {
        event.preventDefault();
        wallets.innerHTML = '';

        fetch('http://localhost:8000/get-wallets')
            .then(response => response.json())
            .then(data => {
                wallets.innerHTML = JSON.stringify(data, null, 4);
            })
            .catch(error => {
                console.error('Произошла ошибка:', error);
            });

    }

    function checkBC(event) {
        event.preventDefault();
        blockChain.innerHTML = '';

        fetch('http://localhost:8000/check-bc-validity')
            .then(response => response.json())
            .then(data => {
                blockChain.innerHTML = JSON.stringify(data, null, 4);
            })
            .catch(error => {
                console.error('Произошла ошибка:', error);
            });

    }

    const transactionForm = document.getElementById('transactionForm');
    transactionForm.addEventListener('submit', sendTransaction);

    const checkStatusTransactionForm = document.getElementById('checkStatusTransactionForm');
    checkStatusTransactionForm.addEventListener('submit', checkStatusTransaction);

    const mineForm = document.getElementById('mineForm');
    mineForm.addEventListener('submit', mine);

    const getWalletsButton = document.getElementById('getWallets');
    getWalletsButton.addEventListener('click', getWallets);

    const checkBCButton = document.getElementById('checkBC');
    checkBCButton.addEventListener('click', checkBC);
</script>
</body>
</html>
`
