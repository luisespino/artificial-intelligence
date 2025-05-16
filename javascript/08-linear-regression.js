async function fit_predict_draw() {
    const { LinearRegression, joinArrays } = await import('https://luisespino.github.io/mlearnjs/mlearn.mjs');

    const myLinearRegression = await LinearRegression(); 
    const model = new myLinearRegression();

    const X = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
    const y = [1, 4, 1, 5, 3, 7, 2, 7, 4, 9];

    model.fit(X, y);

    yPredict = model.predict(X)

    const myjoinArrays = await joinArrays();
    const arr = myjoinArrays('x', X, 'y', y, 'yPredict', yPredict);

    const log = document.getElementById('log');
    const yPred = yPredict.map(num => parseFloat(num.toFixed(2)));
    const mse = model.mse(y, yPredict);
    const r2 = model.r2(y, yPredict);
    log.innerHTML = 'X: '+X+'<br>y: '+y+'<br>yPredict: '+yPred;
    log.innerHTML += '<br>MSE: '+mse+'<br>R2: '+r2;

    google.charts.load('current', {'packages':['corechart']});
    google.charts.setOnLoadCallback(drawChart);    
    function drawChart() {
        var data = google.visualization.arrayToDataTable(arr);
        var options = {
            series: {
                0: {type: 'scatter'},
                1: {type: 'line'}}            
        };  
        var chart = new google.visualization.ComboChart(document.getElementById('chart_div'));
        chart.draw(data, options);         
    }
}

fit_predict_draw();
