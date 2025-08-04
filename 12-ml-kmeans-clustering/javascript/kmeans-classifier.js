

import { KNearestNeighbors, accuracyScore } from 'https://snugml.github.io/js/ml.mjs';

const X = [
    [1, 6],
    [2, 7],
    [3, 6],
    [4, 2],
    [5, 3],
    [6, 1],
    [7, 5],
    [8, 6],
    [9, 4],
    [10, 5]
];

const y = [0, 0, 0, 2, 2, 2, 1, 1, 1, 1 ];

const table = X.map((fila, i) => [...fila, y[i]]);

showTable(table);

const model = new KNearestNeighbors();
model.fit(X, y);

const yPredict = model.predict(X);

const accuracy = accuracyScore(y, yPredict);

const log = document.getElementById('log');
log.innerHTML += '<br><br>Predict:<br>'+ JSON.stringify(yPredict, null, 2);
log.innerHTML += '<br><br>AccuracyScore: '+accuracy;


function showTable(table) {
    let container = document.getElementById('table-container');

    // Crear el elemento de la tabla
    let tableElement = document.createElement('table');

    // Crear la cabecera de la tabla
    let header = tableElement.createTHead();
    let headerRow = header.insertRow();
    let headers = ['X', 'Y', 'CLASS'];
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