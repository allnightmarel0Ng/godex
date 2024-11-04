document.addEventListener('DOMContentLoaded', () => {
    const storeForm = document.getElementById('storeForm');
    const findForm = document.getElementById('findForm');
    const storeResult = document.getElementById('storeResult');
    const findResult = document.getElementById('findResult');

    function getEnvPort() {
        const xhr = new XMLHttpRequest();
        xhr.open('GET', '../.env', false);
        xhr.send(null);

        if (xhr.status === 200) {
            const envContent = xhr.responseText;
            const portMatch = envContent.match(/GATEWAY_PORT=(\d+)/);
            if (portMatch) {
                return portMatch[1];
            }
        }

        return '8080';
    }

    const apiPort = getEnvPort()

    storeForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const link = document.getElementById('link').value;
        try {
            const response = await fetch(`http://localhost:${apiPort}/store`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ link })
            });
            const result = await response.json();
            storeResult.textContent = JSON.stringify(result, null, 2);
        } catch (error) {
            displayError({ message: 'Error storing link: ' + error.message });
        }
    });

    findForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const signature = document.getElementById('signature').value;
        try {
            const response = await fetch(`http://localhost:${apiPort}/find`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ signature })
            });

            if (response.ok) {
                const result = await response.json();
                if (Array.isArray(result)) {
                    displayResults(result);
                } else {
                    displayError({ message: 'Unexpected response format' });
                }
            } else {
                const error = await response.json();
                displayError(error);
            }
        } catch (error) {
            displayError({ message: 'Error finding signature: ' + error.message });
        }
    });

    function displayResults(data) {
        const headers = ['Name', 'Signature', 'Comment', 'File Name', 'Package Name', 'Package Link'];
        const headerRow = headers.map(header => `<th>${header}</th>`).join('');

        const rows = data.map(item => {
            return `
                <tr>
                    <td>${item.name}</td>
                    <td>${item.signature}</td>
                    <td>${item.comment}</td>
                    <td>${item.file.name}</td>
                    <td>${item.file.package.name}</td>
                    <td><a href="${item.file.package.link}" target="_blank">${item.file.package.link}</a></td>
                </tr>
            `;
        }).join('');

        const table = `
            <table class="result-table">
                <thead>
                    <tr>${headerRow}</tr>
                </thead>
                <tbody>
                    ${rows}
                </tbody>
            </table>
        `;

        findResult.innerHTML = table;
    }

    function displayError(error) {
        const errorMessage = document.createElement('div');
        errorMessage.classList.add('error-message');
        errorMessage.textContent = `Error: ${error.message}`;
        findResult.innerHTML = '';
        findResult.appendChild(errorMessage);
    }
});