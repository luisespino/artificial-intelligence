// MIT License
// Copyright (c) 2021 Luis Espino

function heuristic(start, end, h) {
	if (h == 1)	{ // tiles out is sometimes encycled
		var tiles_out = 0
		for (var i = 0; i < start.length; i++){
			if (start[i] != end[i]) tiles_out++
		}
		return tiles_out	
	} else if (h == 2) { // Manhattan
		var man = 0
		for (var i = 0; i < start.length; i++){
			man += Math.abs(i - end.indexOf(start.substring(i,i+1)))
		}
		return man
	}
}


function successors(n, e, h){
	var suc = []
	for (var i = 0; i < n[0].length - 1; i++) {
		let tmp = n[0].substring(i,i+1)
		let child = n[0].substring(0,i)+n[0].substring(i+1,i+2)+tmp+n[0].substring(i+2)
		suc.push([child,heuristic(child, e, h),inc()])
	}
	return suc
}

function bestfirst(start, end, h){
	var cont = 0
	var dot = '{'
	var list = [[start,heuristic(start, end, h),inc()]];
	dot+=list[0][2]+' [label="'+list[0][0]+'"];'
	while (list.length > 0){		
		var current = list.shift();
		if (current[0] == end) {			
			dot += '}'
			return dot
		}		
		var temp = successors(current, end, h);
		//temp.reverse();
		temp.forEach(val => dot+=val[2]+' [label="'+val[0]+'"];'+current[2]+'--'+val[2]+' [label="'+val[1]+'"] ;')
		list = list.concat(temp);
		list = list.sort( function(a,b) { return a[1] - b[1] });
		cont++
		if (cont > 100) {
			alert("The search is looped!")
			dot += '}'
			return dot
		}
	}
	dot += '}'
	return dot
}

var id = 1
function inc() {
	return id++
}

function puzzle() {
	var nodes = prompt("Ingrese texto inicial, texto final y heurstica (1 o 2) separados por un espacio")
	if (nodes == null || nodes == '') nodes = 'halo hola 1'
	nodes = nodes.split(' ')
	return bestfirst(nodes[0], nodes[1], nodes[2])
}
