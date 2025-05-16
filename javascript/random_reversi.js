// obtiene turno oponente
function op(turno) {
    if (turno == 0) return 1
    return 0
}

// carga parametros
var urlParams = new URLSearchParams(window.location.search);
var turno = urlParams.get('turno')
var lineal = urlParams.get('estado')
console.log(turno+' '+lineal)
var matriz = []
var mov = []
var c = 0

// transformación lineal a matriz
for (let i = 0; i < 8; i++) {
    matriz[i] = []
    for (let j = 0; j < 8; j++)
        matriz[i][j] = lineal[c++]
}

//busqueda de movimientos
for (var i = 0; i < 8; i++)
    for (var j = 0; j < 8; j++)
        if (matriz[i][j] == turno){
            //norte
            if (i > 1)
                if (matriz[i-1][j] == op(turno))
                    for (var a = i - 2; a >= 0; a--)
                        if (matriz[a][j] == 2){
                            mov.push(''+a+j)
                            break
                        }
                        else if (matriz[a][j] == turno)
                            break
            //sur
            if (i < 6)
                if (matriz[i+1][j] == op(turno))
                    for (var a=i+2; a <= 7; a++)
                        if (matriz[a][j] == 2){
                            mov.push(''+a+j)
                            break
                        }
                        else if (matriz[a][j] == turno)
                            break
            //oeste
            if (j > 1)
                if (matriz[i][j-1] == op(turno))
                    for (var a = j-2; a >= 0; a--)
                        if (matriz[i][a] == 2){
                            mov.push(''+i+a)
                            break
                        }
                        else if (matriz[i][a] == turno)
                            break;

            //este
            if (j < 6)
                if (matriz[i][j+1] == op(turno))
                    for (var a=j+2; a<=7; a++)
                        if (matriz[i][a] == 2){
                            mov.push(''+i+a)
                            break
                        }
                        else if (matriz[i][a] == turno)
                            break 
            //noroeste
            if (i>1 && j>1)
                if (matriz[i-1][j-1] == op(turno))
                    for (var a=i-2, b=j-2; a>=0,b>=0; a--,b--)
                        if (matriz[a][b] == 2){
                            mov.push(''+a+b)
                            break
                        }
                        else if (matriz[a][b] == turno)
                            break

            //noreste
            if (i>1 && j<6)
                if (matriz[i-1][j+1] == op(turno))
                    for (var a=i-2,b=j+2; a>=0,b<=7; a--,b++)
                        if (matriz[a][b]==2){
                            mov.push(''+a+b)
                            break
                        }
                        else if (matriz[a][b] == turno)
                            break

            //suroeste
            if (i<6 && j>1)
                if (matriz[i+1][j-1] == op(turno))
                    for (a=i+2,b=j-2; a<=7,b>=0; a++,b--)
                        if (matriz[a][b]==2){
                            mov.push(''+a+b)
                            break
                        }
                        else if (matriz[a][b] == turno)
                            break

           //sureste
            if (i<6 && j<6)
                if (matriz[i+1][j+1] == op(turno))
                    for (a=i+2,b=j+2; a<=7,b<=7; a++,b++)
                        if (matriz[a][b] == 2){
                            mov.push(''+a+b)
                            break
                        }
                        else if (matriz[a][b] == turno)
                            break           
        }

//seleccion movimiento mediante random        
document.getElementById("log").innerHTML = mov[Math.floor(Math.random() * mov.length)]