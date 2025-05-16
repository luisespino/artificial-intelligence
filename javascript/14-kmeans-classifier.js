async function fit_predict() {
    const { KMeans } = await import('https://luisespino.github.io/mlearnjs/mlearn.mjs');

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

    showTable(X);

    const myKMeans = await KMeans(); 
    const model = new myKMeans(3, 300, 1e-4, 0);

    
    model.fit(X);

    const yPredict = model.predict(X);

    const centroids = model.getCentroids();
    
    const log = document.getElementById('log');
    log.innerHTML += '<br><br>Predict:<br>'+ JSON.stringify(yPredict, null, 2);

    log.innerHTML += '<br><br>Centroids:<br>'+ JSON.stringify(centroids, null, 2);

    
      
    
    
    google.charts.load('current', { packages: ['corechart'] });
    google.charts.setOnLoadCallback(drawChart);

    function drawChart() {
    const data = new google.visualization.DataTable();
    data.addColumn('number', 'X');
    data.addColumn('number', 'Y');
    data.addColumn({ type: 'string', role: 'style' });


    const colores = ['red', 'green', 'blue']; // Colores por cluster

    // Armar los datos con colores
    const datosGraficados = X.map((p, i) => [p[0], p[1], `point { fill-color: ${colores[yPredict[i]]}; }`]);

    centroids.forEach(c => datosGraficados.push([c[0], c[1], 'point { fill-color: black; }']));

    data.addRows(datosGraficados);

    const options = {
        title: 'Clustering con KMeans',
        hAxis: { title: 'X' },
        vAxis: { title: 'Y' },
        legend: 'none',
        pointSize: 10
    };

    const chart = new google.visualization.ScatterChart(document.getElementById('chart_div'));
    chart.draw(data, options);
    }



}

function showTable(table) {
    let container = document.getElementById('table-container');

    // Crear el elemento de la tabla
    let tableElement = document.createElement('table');

    // Crear la cabecera de la tabla
    let header = tableElement.createTHead();
    let headerRow = header.insertRow();
    let headers = ['X', 'Y'];
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
