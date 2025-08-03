// MIT License
// Copyright (c) 2021 Luis Espino

// revolve Fisher-Yates
function revolve(arr) {
    for (let i = arr.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [arr[i], arr[j]] = [arr[j], arr[i]];
      }
      return arr;
}

function heuristic(start, end) {
	var tiles_out = 0
	for (var i = 0; i < start.length; i++){
		if (start[i] != end[i]) tiles_out++
	}
	return tiles_out	
}

function successors(n, e){
	var suc = []
	for (var i = 0; i < n[0].length - 1; i++) {
        let level = n[3]+1
		let tmp = n[0].substring(i,i+1)
		let child = n[0].substring(0,i)+n[0].substring(i+1,i+2)+tmp+n[0].substring(i+2)
		suc.push([child,heuristic(child, e)+level,inc(),level]) 
	}
	return suc
}

function genetic(end){
    // define number of individuals, length and max value
    var individuals = 4
    var len = end.length
    var max = ''
    for (var i = 0; i < len; i++) {
        max += '9'
    }
    max = parseInt(max)
    end = parseInt(end)
    document.getElementById("log").innerHTML+="Objetive: "+end
    
    // initialize population
    var population = []
    for (var i = 0; i < individuals; i++){
        population[i] = Math.floor(Math.random() * max + 1)
    }
    
    var expected = max
    var itera = 0
    while (expected > 1){
        itera++
        document.getElementById("log").innerHTML+="<br><br>Population "+itera+": "+population
        
        // static selection by tournament
        if (Math.abs(end-population[0]) <= Math.abs(end-population[1])) p1 = 0
        else p1 = 1
        if (Math.abs(end-population[2]) <= Math.abs(end-population[3])) p2 = 2
        else p2 = 3
        document.getElementById("log").innerHTML+="<br>Parents selected: "+population[p1]+' '+population[p2]


        // crossover
        let child = []
        child.push(Math.floor((population[p1]+population[p2])/2))
        child.push(Math.floor(Math.abs(2*population[p1]-population[p2])))
        child.push(Math.floor(Math.abs(population[p1]-population[p2])))
        child.push(Math.floor(population[p1]*1.1))
        child.push(Math.floor(population[p1]*.9))
        child.push(Math.floor(population[p1]*.5))
        document.getElementById("log").innerHTML+="<br>Crossover: "+child

        // optional mutation 

        // evaluation
        if (Math.abs(end-child[2]) <= Math.abs(end-child[0])) child[0] = child[2]
        if (Math.abs(end-child[3]) <= Math.abs(end-child[1])) child[1] = child[3]
        if (Math.abs(end-child[4]) <= Math.abs(end-child[0])) child[0] = child[4]
        if (Math.abs(end-child[5]) <= Math.abs(end-child[1])) child[1] = child[5]

        // replacement and shuffle
        population[0] = population[p1]
        population[1] = child[0]
        population[2] = population[p2]
        population[3] = child[1]
        population = revolve(population)
        document.getElementById("log").innerHTML+="<br>Replacement: "+population
        const sum = population.reduce((a, b) => Math.abs(end-a) + Math.abs(end-b), 0);
        //const avg = (sum / population.length) || 0;
        expected = sum //Math.abs(end-avg)
        document.getElementById("log").innerHTML+="<br>AVG Expected value: "+expected
        if (itera>=100) {
            document.getElementById("log").innerHTML+="<br><br>The algorithm is looped!"
            break 
        } 
        
    }

    /*
	var cont = 0
	var dot = '{'
	var list = [[start,heuristic(start, end),inc(),0]];
	dot+=list[0][2]+' [label="'+list[0][0]+'"];'
	while (list.length > 0){		
		var current = list.shift();
		if (current[0] == end) {			
			dot += '}'
			return dot
		}		
		var temp = successors(current, end);
		//temp.reverse();
		temp.forEach(val => dot+=val[2]+' [label="'+val[0]+'"];'+current[2]+'--'+val[2]+' [label="'+(val[1]-current[3]-1)+'+'+(current[3]+1)+'"] ;')
		list = list.concat(temp);
		list = list.sort( function(a,b) { return a[1] - b[1] });
		cont++
		if (cont > 100) {
			alert("The search is looped!")
			dot += '}'
			return dot
		}
	}
    */

}

var num = prompt("Enter an integer to search for:")
if (num == null || num == '') num = '1234'
genetic(num)
