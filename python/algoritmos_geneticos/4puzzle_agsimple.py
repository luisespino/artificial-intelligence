import random

def ag(n,i):
    gen = []

    #generar primera generación
    for x in range(0,4):
        gen.extend([random.randint(0,9999)])
    print(gen)  

    #iterar i veces el algoritmo genético
    for x in range(0, i):

        #selección de 2 padres
        if (abs(n-gen[0])<=abs(n-gen[1])): p1 = 0
        else: p1 = 1
        if (abs(n-gen[2])<=abs(n-gen[3])): p2 = 2
        else: p2 = 3

        #generación de 4 hijos y selección de 2
        h1 = int((gen[p1]+gen[p2])/2)
        h2 = abs(2*gen[p1]-gen[p2])
        h3 = abs(gen[p1]-gen[p2])
        h4 = int(gen[p1]*1.2)
        #mutación
        #h4 = random.randint(0,9999)

        if (abs(n-h4)<=abs(n-h1)):
           h1 = h4        
        if (abs(n-h3)<=abs(n-h2)):
           h2 = h3

        #reemplazo
        gen[0] = gen[p1]
        gen[1] = gen[p2]
        gen[2] = h1
        gen[3] = h2
        random.shuffle(gen)
        print(gen)
 
ag(1234,15)

