# Luis Espino 2020

import matplotlib.pyplot as plt
from sklearn.tree import DecisionTreeClassifier, plot_tree

# Data creation
outlook=['sunny','sunny','overcast','rain','rain','rain','overcast',
         'sunny','sunny','rain','sunny','overcast','overcast','rain']
temperature=['hot','hot','hot','mild','cool','cool','cool',
             'mild','cool','mild','mild','mild','hot','mild']
humidity=['high','high','high','high','normal','normal','normal',
          'high','normal','normal','normal','high','normal','high']
windy=['false','true','false','false','false','true','true',
       'false','false','false','true','true','false','true']
play=['N','N','P','P','P','N','P','N','P','P','P','P','P','N']

# Import LabelEncoder
from sklearn import preprocessing

# Creating labelEncoder
le = preprocessing.LabelEncoder()

# Converting string labels into numbers.
outlook_encoded=le.fit_transform(outlook)
temperature_encoded=le.fit_transform(temperature)
humidity_encoded=le.fit_transform(humidity)
windy_encoded=le.fit_transform(windy)
label=le.fit_transform(play)

print ("outlook:  ",outlook_encoded)
print ("temp:     ",temperature_encoded)
print ("humidity: ",humidity_encoded)
print ("windy:    ",windy_encoded)
print ("PLAY:     ",label)

# Combinig attributes into single listof tuples
features=list(zip(outlook_encoded,temperature_encoded,humidity_encoded,windy_encoded))
print (features)

# fit the model
clf = DecisionTreeClassifier().fit(features, label)
plot_tree(clf, filled=True)
plt.show()

