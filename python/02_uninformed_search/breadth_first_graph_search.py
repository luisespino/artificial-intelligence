# Luis Espino 2018
#
# 4 Lineal Puzzle
# the successors are the left, center and right exchanges
# until a solution is found

def succesors(n):
    return [n[1]+n[0]+n[2]+n[3],n[0]+n[2]+n[1]+n[3],n[0]+n[1]+n[3]+n[2]]

def breadth_first_graph_search (begin, end):
    no_visited = [begin]
    visited = []
    while no_visited:
        current = no_visited.pop(0)
        visited.append(current)
        print ("Current: "+current)
        if current == end:
            print (str(len(no_visited)+len(visited))+" nodes")
            return print ("SOLUTION")
        temp = succesors (current)
        print ("Succesors: "+str(temp))
        if temp:
            a = no_visited + visited
            no_visited.extend([x for x in temp if x not in a])
            print ("No visited: "+str(no_visited))
    print ("NO SOLUTION")

breadth_first_graph_search ("1423","1234")
