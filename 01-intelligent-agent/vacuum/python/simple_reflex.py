# Luis Espino 2017

import time

def reflex_agent(location, state):
    if state=="DIRTY":
        return 'CLEAN'
    elif location=='A':
        return 'RIGHT'
    elif location=='B':
        return 'LEFT'   

def test(states):
    while True:
        location = states[0]
        state = (states[2], states[1])[states[0]=='A']
        action = reflex_agent(location, state)
        print ("Location: "+location+" | Action: "+action)
        if action == "CLEAN":
            if location == 'A':
                states[1]="CLEAN"
            elif location == 'B':
                states[2]="CLEAN"
        elif action == "RIGHT":
            states[0]='B'
        elif action == "LEFT":
            states[0]='A' 
        time.sleep(3)

test(['A','DIRTY','DIRTY'])
