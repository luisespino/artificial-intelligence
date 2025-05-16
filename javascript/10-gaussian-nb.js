async function fit_predict() {
    const { GaussianNB, LabelEncoder, accuracyScore } = await import('https://luisespino.github.io/mlearnjs/mlearn.mjs');

    

    const outlook = ['sunny', 'sunny', 'overcast', 'rain', 'rain', 'rain', 'overcast',
        'sunny', 'sunny', 'rain', 'sunny', 'overcast', 'overcast', 'rain'];

    const temperature = ['hot', 'hot', 'hot', 'mild', 'cool', 'cool', 'cool',
        'mild', 'cool', 'mild', 'mild', 'mild', 'hot', 'mild'];

    const humidity = ['high', 'high', 'high', 'high', 'normal', 'normal', 'normal',
        'high', 'normal', 'normal', 'normal', 'high', 'normal', 'high'];

    const windy = ['false', 'true', 'false', 'false', 'false', 'true', 'true',
        'false', 'false', 'false', 'true', 'true', 'false', 'true'];

    const label = ['N', 'N', 'P', 'P', 'P', 'N', 'P', 'N', 'P', 'P', 'P', 'P', 'P', 'N'];

    const table = outlook.map((_, i) => [outlook[i], temperature[i], humidity[i], windy[i], label[i]]);

    showTable(table);

    
    const myLabelEncoder = await LabelEncoder(); 
    const encoder = new myLabelEncoder();

    const encOut = encoder.fitTransform(outlook);
    const encTem = encoder.fitTransform(temperature);
    const encHum = encoder.fitTransform(humidity);
    const encWin = encoder.fitTransform(windy);
    const encLab = encoder.fitTransform(label);

    const features = encOut.map((_, i) => [encOut[i], encTem[i], encHum[i], encWin[i]]);
     

    const myGaussianNB = await GaussianNB(); 
    const model = new myGaussianNB();

    model.fit(features, encLab);

    const encYPredict = model.predict(features)
    const yPredict = encoder.inverseTransform(encYPredict);

    const myAccuracyScore = await accuracyScore();
    const accuracy = myAccuracyScore(encLab, encYPredict);

    const log = document.getElementById('log');
    log.innerHTML = '<br><br>LabelEncoder:<br>'+JSON.stringify(features, null, 2);
    log.innerHTML += '<br><br>Predict:<br>'+ JSON.stringify(yPredict, null, 2);
    log.innerHTML += '<br><br>AccuracyScore: '+accuracy;
}

function showTable(table) {
    let container = document.getElementById('table-container');

    // Crear el elemento de la tabla
    let tableElement = document.createElement('table');

    // Crear la cabecera de la tabla
    let header = tableElement.createTHead();
    let headerRow = header.insertRow();
    let headers = ['Outlook', 'Temp', 'Humidity', 'Windy', 'Label'];
    headers.forEach(headerText => {
        let cell = headerRow.insertCell();
        cell.textContent = headerText;
    });

    // Crear el cuerpo de la tabla
    let body = tableElement.createTBody();
    table.forEach(rowData => {
        let row = body.insertRow();
        rowData.forEach(cellData => {
            let cell = row.insertCell();
            cell.textContent = cellData;
        });
    });

    // Insertar la tabla en el contenedor
    container.appendChild(tableElement);
}

fit_predict();

