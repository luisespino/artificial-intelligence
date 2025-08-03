#Autor: Luis Espino 2016

def heuristic(nodo_actual,nodo_fin):
    return [x != y for (x, y) in zip(nodo_actual, nodo_fin)].count(True)

def sucesores(n,s):
    return [[n[1]+n[0]+n[2]+n[3],heuristic(n[1]+n[0]+n[2]+n[3],s)],
            [n[0]+n[2]+n[1]+n[3],heuristic(n[0]+n[2]+n[1]+n[3],s)],
            [n[0]+n[1]+n[3]+n[2],heuristic(n[0]+n[1]+n[3]+n[2],s)]]

def bestfirst (nodo_inicio, nodo_fin):
    print ("\nBEST-FIRST")
    lista = [[nodo_inicio, heuristic(nodo_inicio, nodo_fin)]]
    while lista:
        nodo_actual = lista.pop(0)
        print (nodo_actual)
        if nodo_actual[0] == nodo_fin:
            return print ("SOLUCION")
        temp = sucesores (nodo_actual[0],nodo_fin)
        print (temp)
        if temp:
            lista.extend(temp)
            lista.sort(key=lambda x: int(x[1]))
            print (lista)
    print ("NO-SOLUCION")

bestfirst("halo","hola")
