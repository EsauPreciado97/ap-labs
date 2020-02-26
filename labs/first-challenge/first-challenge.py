str = input("Input: ")

longest = ''
index = 0

 
while(str[index] not in longest and index < len(str)):
    longest += str[index]
    index += 1

another = ''

while(index < len(str)):
    if(str[index] in another):
        another = str[index]
    else:
        another += str[index]
    if len(another) > len(longest):
        longest = another
        another = ''
    index += 1


print('Output: ',len(longest))  


    



        
