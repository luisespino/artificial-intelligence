# Luis Espino 2017
import random

# tiles out heuristic
def h(current, end):
    return [x != y for (x, y) in zip(current, end)].count(True)

def succesors(n,s):
    i = n.index('0')
    if i==0: return   [[n[1]+n[0]+n[2]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8], h(n[1]+n[0]+n[2]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[3]+n[1]+n[2]+n[0]+n[4]+n[5]+n[6]+n[7]+n[8], h(n[3]+n[1]+n[2]+n[0]+n[4]+n[5]+n[6]+n[7]+n[8],s)]]
    elif i==1: return [[n[1]+n[0]+n[2]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8], h(n[1]+n[0]+n[2]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[2]+n[1]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8], h(n[0]+n[2]+n[1]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[4]+n[2]+n[3]+n[1]+n[5]+n[6]+n[7]+n[8], h(n[0]+n[4]+n[2]+n[3]+n[1]+n[5]+n[6]+n[7]+n[8],s)]]
    elif i==2: return [[n[0]+n[2]+n[1]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8], h(n[0]+n[2]+n[1]+n[3]+n[4]+n[5]+n[6]+n[7]+n[8],s)],                       
                       [n[0]+n[1]+n[5]+n[3]+n[4]+n[2]+n[6]+n[7]+n[8], h(n[0]+n[1]+n[5]+n[3]+n[4]+n[2]+n[6]+n[7]+n[8],s)]]
    elif i==3: return [[n[3]+n[1]+n[2]+n[0]+n[4]+n[5]+n[6]+n[7]+n[8], h(n[3]+n[1]+n[2]+n[0]+n[4]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[4]+n[3]+n[5]+n[6]+n[7]+n[8], h(n[0]+n[1]+n[2]+n[4]+n[3]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[6]+n[4]+n[5]+n[3]+n[7]+n[8], h(n[0]+n[1]+n[2]+n[6]+n[4]+n[5]+n[3]+n[7]+n[8],s)]]
    elif i==4: return [[n[0]+n[4]+n[2]+n[3]+n[1]+n[5]+n[6]+n[7]+n[8], h(n[0]+n[4]+n[2]+n[3]+n[1]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[4]+n[3]+n[5]+n[6]+n[7]+n[8], h(n[0]+n[1]+n[2]+n[4]+n[3]+n[5]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[5]+n[4]+n[6]+n[7]+n[8], h(n[0]+n[1]+n[2]+n[3]+n[5]+n[4]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[7]+n[5]+n[6]+n[4]+n[8], h(n[0]+n[1]+n[2]+n[3]+n[7]+n[5]+n[6]+n[4]+n[8],s)]]
    elif i==5: return [[n[0]+n[1]+n[5]+n[3]+n[4]+n[2]+n[6]+n[7]+n[8], h(n[0]+n[1]+n[5]+n[3]+n[4]+n[2]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[5]+n[4]+n[6]+n[7]+n[8], h(n[0]+n[1]+n[2]+n[3]+n[5]+n[4]+n[6]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[4]+n[8]+n[6]+n[7]+n[5], h(n[0]+n[1]+n[2]+n[3]+n[4]+n[8]+n[6]+n[7]+n[5],s)]]    
    elif i==6: return [[n[0]+n[1]+n[2]+n[6]+n[4]+n[5]+n[3]+n[7]+n[8], h(n[0]+n[1]+n[2]+n[6]+n[4]+n[5]+n[3]+n[7]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[7]+n[6]+n[8], h(n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[7]+n[6]+n[8],s)]]    
    elif i==7: return [[n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[7]+n[6]+n[8], h(n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[7]+n[6]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[7]+n[5]+n[6]+n[4]+n[8], h(n[0]+n[1]+n[2]+n[3]+n[7]+n[5]+n[6]+n[4]+n[8],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[6]+n[8]+n[7], h(n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[6]+n[8]+n[7],s)]]    
    elif i==8: return [[n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[6]+n[8]+n[7], h(n[0]+n[1]+n[2]+n[3]+n[4]+n[5]+n[6]+n[8]+n[7],s)],
                       [n[0]+n[1]+n[2]+n[3]+n[4]+n[8]+n[6]+n[7]+n[5], h(n[0]+n[1]+n[2]+n[3]+n[4]+n[8]+n[6]+n[7]+n[5],s)]]

def mix(current, end):
    return random.choice(succesors(current, end))[0]

def bestfirst (end, n):
    # mix the end node
    start = end
    for x in range (0, n):
        start = mix(start, end)
    
    # start the search 
    print ("\nBEST-FIRST")
    unvisited = [[start, h(start, end)]]
    visited = []
    count = 0
    while unvisited:
        count = count + 1
        current = unvisited.pop(0)        
        print(current)
        visited.extend([current[0]])
        if current[0] == end:
            return print("SOLUTION IN "+str(count)+" ITERATIONS")
        temp = succesors(current[0], end)
        print(temp)
        # without repeating visited nodes
        unvisited.extend([x for x in temp if x[0] not in visited])
        unvisited.sort(key=lambda x: int(x[1]))
        print (unvisited)
    print ("NO-SOLUTION")

# bestfirst (end = 123456780, mix_mov = 20)
bestfirst("123456780", 20)
