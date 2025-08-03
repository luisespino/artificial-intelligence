def succesor(n):
    if n[0] == 'A': return [['B'], ['C'], ['D']]
    elif n[0] == 'B': return [['C'], ['E']]
    elif n[0] == 'C': return [['E'], ['F'], ['G']]
    elif n[0] == 'D': return [['C'], ['F']]
    elif n[0] == 'E': return [['G']]
    elif n[0] == 'F': return [['G']]
    else: return None


def museum (begin, end):
    numsol = 0
    novisited = [[begin]]
    while novisited:
        current = novisited.pop(0)
        print (current)
        if current[0] == end:
            numsol += 1
            # show inverse order 
            print (current[:][::-1])
            print("SOLUTION")
        temp = succesor(current)
        # optional: temp.reverse()
        print (temp)
        if temp:
            for x in temp:
                x.extend(current)
            novisited.extend(temp)
            print(novisited)
    print("TOTAL: "+str(numsol))

museum('A','G')
