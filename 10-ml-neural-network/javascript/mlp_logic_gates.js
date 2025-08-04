
import { MLPClassifier, LabelEncoder, accuracyScore } from 'https://snugml.github.io/js/ml.mjs';

const X = [
    [0, 0],
    [0, 1],
    [1, 0],
    [1, 1]
];

const y1 = [0, 1, 1, 1];
const y2 = [0, 0, 0, 1];
const y3 = [0, 1, 1, 0];

const model = new MLPClassifier(2, 4, 1);
model.fit(X, y1);
let yPredict = X.map(f => {
    const output = model.predict(f);
    if (output.length === 1) {
        // Binaria: redondear
        return Math.round(output[0]);
    } else {
        // Multiclase: elegir la clase con mayor probabilidad
        return output.indexOf(Math.max(...output));
    }
    });

let accuracy = accuracyScore(y1, yPredict);
const log = document.getElementById('log');
log.innerHTML = '<br><br>OR: X = '+JSON.stringify(X, null, 2);
log.innerHTML += '; y = '+JSON.stringify(y1, null, 2);
log.innerHTML += '; yPredict = '+ JSON.stringify(yPredict, null, 2);
log.innerHTML += '; AccuracyScore: '+accuracy;


model.fit(X, y2);
yPredict = X.map(f => {
    const output = model.predict(f);
    if (output.length === 1) {
        // Binaria: redondear
        return Math.round(output[0]);
    } else {
        // Multiclase: elegir la clase con mayor probabilidad
        return output.indexOf(Math.max(...output));
    }
    });

accuracy = accuracyScore(y2, yPredict);
log.innerHTML += '<br><br>AND: X = '+JSON.stringify(X, null, 2);
log.innerHTML += '; y = '+JSON.stringify(y2, null, 2);
log.innerHTML += '; yPredict = '+ JSON.stringify(yPredict, null, 2);
log.innerHTML += '; AccuracyScore: '+accuracy;


model.fit(X, y3);
yPredict = X.map(f => {
    const output = model.predict(f);
    if (output.length === 1) {
        // Binaria: redondear
        return Math.round(output[0]);
    } else {
        // Multiclase: elegir la clase con mayor probabilidad
        return output.indexOf(Math.max(...output));
    }
    });

accuracy = accuracyScore(y3, yPredict);
log.innerHTML += '<br><br>XOR: X = '+JSON.stringify(X, null, 2);
log.innerHTML += '; y = '+JSON.stringify(y3, null, 2);
log.innerHTML += '; yPredict = '+ JSON.stringify(yPredict, null, 2);
log.innerHTML += '; AccuracyScore: '+accuracy;


const modelXOR = new MLPClassifier(2, 5, 1);

modelXOR.fit(X, y3);
yPredict = X.map(f => {
    const output = modelXOR.predict(f);
    if (output.length === 1) {
        // Binaria: redondear
        return Math.round(output[0]);
    } else {
        // Multiclase: elegir la clase con mayor probabilidad
        return output.indexOf(Math.max(...output));
    }
    });

accuracy = accuracyScore(y3, yPredict);
log.innerHTML += '<br><br>XOR (extra-neuron): X = '+JSON.stringify(X, null, 2);
log.innerHTML += '; y = '+JSON.stringify(y3, null, 2);
log.innerHTML += '; yPredict = '+ JSON.stringify(yPredict, null, 2);
log.innerHTML += '; AccuracyScore: '+accuracy;    